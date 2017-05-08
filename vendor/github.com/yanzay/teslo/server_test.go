package teslo

import (
	"sync"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	if s == nil {
		t.Error("NewServer should initialize Server")
	}
	if s.Addr == "" {
		t.Error("Default address should not be empty")
	}
}

func TestStart(t *testing.T) {
	s := NewServer()
	go s.Start()
}

func TestRespond(t *testing.T) {
	s := NewServer()
	id := "test"
	content := "content"
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		resp := <-s.responses
		if resp.ID != "test" {
			t.Errorf("Message ID should be %s", id)
		}
		if resp.Content != "content" {
			t.Errorf("Message Content should be %s", content)
		}
		wg.Done()
	}()
	s.Respond(id, content)
	wg.Wait()
}

func TestSubscribe(t *testing.T) {
	s := NewServer()
	s.Subscribe("test", func(Responder, *Event) {})
	if s.handlers["test"] == nil {
		t.Error("Subscribe should add handler to handlers")
	}
}
