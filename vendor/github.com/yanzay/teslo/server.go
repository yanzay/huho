package teslo

import (
	"fmt"
	"io"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
	"github.com/yanzay/log"
)

type Message struct {
	ID      string
	Content string
}

type Responder interface {
	Respond(string, string)
}

type Server struct {
	InitSession  func(string)
	CloseSession func(string)
	Render       func(io.Writer, *http.Request)
	Addr         string
	handlers     map[string]func(*Session, *Event)
}

func NewServer() *Server {
	return &Server{
		Addr:         ":8080",
		handlers:     make(map[string]func(*Session, *Event)),
		InitSession:  func(string) {},
		CloseSession: func(string) {},
	}
}

func (s *Server) Start() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/favicon.ico", ServeFile)
	http.HandleFunc("/", s.IndexHandler)
	http.HandleFunc("/ws", s.WSHandler)
	log.Infof("Starting server at %s", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	http.ServeFile(w, r, "static/"+name)
}

func (s *Server) Subscribe(id string, handler func(*Session, *Event)) {
	s.handlers[id] = handler
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("teslo-session")
	if err != nil || cookie.Value == "" {
		http.SetCookie(w, &http.Cookie{Name: "teslo-session", Value: uuid.NewV4().String()})
	}
	if s.Render != nil {
		s.Render(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Define Render function")
	}
}

func (s *Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("teslo-session")
	if err != nil {
		log.Errorf("Websocket connection without session, terminating")
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Can't update to WebSocket protocol: %s", err)
		return
	}
	if s.InitSession != nil {
		session := NewSession(s, conn, cookie.Value)
		s.InitSession(session.ID)
		session.Start()
	}

}
