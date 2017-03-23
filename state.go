package main

import (
	"fmt"
	"sync"

	"github.com/yanzay/huho/templates"
)

type stateStore struct {
	states map[string]*templates.State
	sync.Mutex
}

func newStateStore() *stateStore {
	return &stateStore{
		states: make(map[string]*templates.State),
	}
}

func (ss *stateStore) GetState(id string) templates.State {
	ss.Lock()
	fmt.Println(ss.states, id)
	state := *ss.states[id]
	ss.Unlock()
	return state
}

func (ss *stateStore) SaveState(id string, state templates.State) {
	ss.Lock()
	fmt.Println("savestate", id, state)
	ss.states[id] = &state
	ss.Unlock()
}

func (ss *stateStore) DeleteState(id string) {
	ss.Lock()
	fmt.Println("deletestate", id)
	delete(ss.states, id)
	ss.Unlock()
}
