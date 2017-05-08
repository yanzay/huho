package main

import (
	"fmt"
	"io"

	"github.com/yanzay/log"
	"github.com/yanzay/teslo"
	"github.com/yanzay/teslo/cmd/templates"
)

var defaultState = templates.State{
	Items: []*templates.Item{
		{ID: "1", Content: "write a todo app"},
		{ID: "2", Content: "extract framework", Done: true},
		{ID: "3", Content: "write a shop app"},
	},
}
var sessions = map[string]templates.State{}

func main() {
	log.Level = log.LevelTrace
	server := teslo.NewServer()
	server.Render = func(w io.Writer) {
		templates.WritePage(w, defaultState)
	}
	server.InitSession = func(id string) {
		var items = []*templates.Item{
			{ID: "1", Content: "write a todo app"},
			{ID: "2", Content: "extract framework", Done: true},
			{ID: "3", Content: "write a shop app"},
		}

		sessions[id] = templates.State{Items: items}
	}
	server.CloseSession = func(id string) {
		fmt.Printf("Closing session with id: %s\n", id)
		delete(sessions, id)
	}
	server.Subscribe("todo", TodoHandler)
	server.Start()
}

func TodoHandler(s *teslo.Session, event *teslo.Event) {
	fmt.Printf("Sessions: %d\n", len(sessions))
	if event.Type == "click" {
		for _, item := range sessions[s.ID].Items {
			if event.ID == item.ID {
				item.Done = !item.Done
			}
		}
		s.Respond("todo", templates.Todo(sessions[s.ID].Items))
		fmt.Printf("Items: %v\n", sessions[s.ID].Items)
	}
}
