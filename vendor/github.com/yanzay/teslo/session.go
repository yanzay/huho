package teslo

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo/templates"
)

var upgrader = websocket.Upgrader{}

type Session struct {
	ID        string
	conn      *websocket.Conn
	requests  chan Event
	responses chan Message
	server    *Server
	close     chan struct{}
}

type Event struct {
	Type    string   `json:"event"`
	ID      string   `json:"id"`
	Parents []string `json:"parents"`
	Data    string   `json:"data"`
}

func NewSession(server *Server, conn *websocket.Conn, sessionID string) *Session {
	log.Debugf("Creating new session %v", conn)
	return &Session{
		ID:        sessionID,
		server:    server,
		conn:      conn,
		requests:  make(chan Event, 0),
		responses: make(chan Message, 0),
		close:     make(chan struct{}, 0),
	}
}

func (s *Session) Start() {
	log.Debugf("Starting session: %s", s.ID)
	defer s.server.CloseSession(s.ID)
	go s.writeLoop()
	s.readLoop()
}

func (s *Session) Close() {
	log.Debugf("Closing session: %s", s.ID)
	close(s.close)
}

func (s *Session) Respond(id string, content string) {
	s.responses <- Message{ID: id, Content: content}
}

func (s *Session) handleRequest(request *Event) {
	log.Debugf("Handling request: %v", request)
	if s.server.handlers[request.ID] != nil {
		log.Debug("Running handler")
		s.server.handlers[request.ID](s, request)
		log.Debug("Handler exited")
	} else {
		for _, parent := range request.Parents {
			if s.server.handlers[parent] != nil {
				s.server.handlers[parent](s, request)
				return
			}
		}
	}
}

func (s *Session) writeLoop() {
	log.Debug("Starting writeLoop")
	for {
		select {
		case <-s.close:
			log.Debugf("Exiting writeLoop for session: %s", s.ID)
			return
		case message := <-s.responses:
			writer, err := s.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Errorf("Can't get next writer for websocket: %s", err)
				s.Close()
				return
			}
			log.Debugf("Sending response: %v", message)
			templates.WriteMessage(writer, message.ID, message.Content)
			err = writer.Close()
			if err != nil {
				log.Errorf("Can't write to websocket: %s", err)
				s.Close()
				return
			}
		}
	}
}

func (s *Session) readLoop() {
	defer log.Debug("exiting readLoop")
	log.Debug("Starting readLoop")
	for {
		log.Debug("Getting next reader")
		_, reader, err := s.conn.NextReader()
		log.Debugf("reader ready, error: %s", err)
		if err != nil {
			log.Errorf("Can't get websocket reader: %s", err)
			log.Debug("Closing session")
			s.Close()
			log.Debug("Session closed")
			return
		}
		decoder := json.NewDecoder(reader)
		e := &Event{}
		decoder.Decode(e)
		s.handleRequest(e)
		log.Debug("request handled")
	}
}
