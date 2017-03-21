package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yanzay/huho/templates"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var defaultState templates.State
var clientID, clientSecret string
var githubConf *oauth2.Config
var store *Storage
var sessions map[string]*templates.State

type UserEmail struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

type Site struct {
	Repo   string `json:"repo"`
	Domain string `json:"domain"`
}

func init() {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		log.Fatal("GITHUB_CLIENT_ID should not be empty")
	}
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("GITHUB_CLIENT_SECRET should not be empty")
	}
	githubConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	url := githubConf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	defaultState.GithubURL = url
	sessions = map[string]*templates.State{"": &defaultState}
}

func main() {
	store = NewStorage()
	defer store.Close()
	server := teslo.NewServer()
	server.InitSession = func(id string) {
		fmt.Println("Session id: ", id)
	}
	server.CloseSession = func(id string) {
		fmt.Println("Closing session:", id)
		delete(sessions, id)
	}
	server.Render = func(w io.Writer, r *http.Request) {
		sessionID := sessionIDFromRequest(r)
		email := store.GetSession(sessionID)
		sessions[sessionID] = &templates.State{GithubURL: defaultState.GithubURL, Login: email}
		templates.WritePage(w, *sessions[sessionID])
	}
	http.HandleFunc("/callback", AuthCallback)
	server.Start()
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	sessionID := sessionIDFromRequest(r)
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	tok, err := githubConf.Exchange(ctx, code)
	if err != nil {
		log.Error(err)
		return
	}

	client := githubConf.Client(ctx, tok)
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		log.Error(err)
		return
	}
	emails := []*UserEmail{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&emails)
	if err != nil {
		log.Error(err)
		return
	}
	for _, email := range emails {
		if email.Primary {
			authorizeUser(sessionID, email.Email)
			http.Redirect(w, r, "/", 302)
		}
	}
}

func sessionIDFromRequest(r *http.Request) string {
	cookie, err := r.Cookie("teslo-session")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func authorizeUser(sessionID, email string) {
	log.Infof("AuthorizeUser", sessionID, email)
	store.StoreSession(sessionID, email)
}
