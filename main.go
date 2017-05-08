package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "net/http/pprof"

	uuid "github.com/satori/go.uuid"
	"github.com/yanzay/huho/templates"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	wwwDir = flag.String("d", "/www", "Directory for generated static files")
)

var (
	defaultState           templates.State
	clientID, clientSecret string
	githubConf             *oauth2.Config
	store                  *Storage
	sessions               = newStateStore()
)

type UserEmail struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
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
}

func main() {
	flag.Parse()
	store = NewStorage()
	defer store.Close()
	server := teslo.NewServer()
	server.InitSession = func(id string) {
		fmt.Println("Session id: ", id)
	}
	server.CloseSession = func(id string) {
		fmt.Println("Closing session:", id)
		//sessions.DeleteState(id)
	}
	server.Render = func(w io.Writer, r *http.Request) {
		sessionID := sessionIDFromRequest(r)
		email := store.GetSession(sessionID)
		projects := store.GetProjects(email)
		sessions.SaveState(sessionID, templates.State{GithubURL: defaultState.GithubURL, Login: email, Projects: projects})
		templates.WritePage(w, sessions.GetState(sessionID))
	}
	http.HandleFunc("/callback", AuthCallback)
	server.Subscribe("addproject", NewProjectHandler)
	server.Subscribe("projects", ChangeProjectHandler)
	server.Start()
}

func NewProjectHandler(session *teslo.Session, event *teslo.Event) {
	fmt.Println("NewProjectHandler")
	if event.Type == "submit" {
		project := parseProject(event.Data)
		fmt.Println("session", sessions.GetState(session.ID))
		fmt.Println("session", session)
		fmt.Println("event", event)
		state := sessions.GetState(session.ID)
		state.Projects = append(state.Projects, project)
		sessions.SaveState(session.ID, state)
		store.StoreProjects(state.Login, state.Projects)
		session.Respond("projects", templates.ProjectList(state))
		session.Respond("addproject", templates.AddProject())
	}
}

func ChangeProjectHandler(session *teslo.Session, event *teslo.Event) {
	fmt.Println("ChangeProjectHandler")
	fmt.Println(event)
	fmt.Println("data", event.Data)
	if event.Type == "change" {
		data := struct {
			ID      string
			Checked bool
		}{}
		json.Unmarshal([]byte(event.Data), &data)
		fmt.Println("parsed data:", data)
		state := sessions.GetState(session.ID)
		for i, project := range state.Projects {
			if project.ID == data.ID {
				state.Projects[i].AutoDeploy = data.Checked
				store.StoreProjects(state.Login, state.Projects)
				sessions.SaveState(session.ID, state)
				return
			}
		}
	}
	if event.Type == "click" {
		if strings.HasPrefix(event.ID, "deploy") {
			id := strings.TrimPrefix(event.ID, "deploy-")
			log.Debug(id)
			state := sessions.GetState(session.ID)
			for _, project := range state.Projects {
				if project.ID == id {
					err := deploy(*wwwDir, project)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}
}

func parseProject(data string) templates.Project {
	vs, err := url.ParseQuery(data)
	if err != nil {
		log.Error(err)
	}
	id := uuid.NewV4().String()
	return templates.Project{
		Domain: vs.Get("domain"),
		URL:    vs.Get("url"),
		ID:     id,
	}
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
