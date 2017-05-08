package main

import (
	"io"

	"github.com/yanzay/teslo"
	"github.com/yanzay/teslo/cmd/templates"
)

func main() {
	InitDB()
	defaultState = templates.State{}
	server := teslo.NewServer()
	server.Render = func(w io.Writer) {
		templates.WritePage(w, faultState)
	}
	server.InitSession = func(id string) {
		products := loadProducts()
		sessions[id] = &templates.tate{Products: products}
	}
	server.CloseSession = func(id string) {
		delete(sessions, id)
	}
	server.Subscribe("login", LoginHandler)
	server.Start()
}

func LoginHandler() {
}
