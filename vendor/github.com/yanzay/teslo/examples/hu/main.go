package main

import (
	"io"

	"github.com/yanzay/teslo"
	"github.com/yanzay/teslo/examples/hu/templates"
)

var defaultState = templates.State{}
var sessions = map[string]*templates.State{}

func main() {
	server := teslo.NewServer()
	server.Render = func(w io.Writer) {
		templates.WritePage(w, defaultState)
	}
	server.InitSession = func(id string) {
		sessions[id] = &templates.State{}
	}
	server.CloseSession = func(id string) {
		delete(sessions, id)
	}
	server.Subscribe("add-repo", AddRepoHandler)
	server.Start()
}

func AddRepoHandler(session *teslo.Session, event *teslo.Event) {

}
