package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/yanzay/huho/templates"
	"github.com/yanzay/log"
)

type Storage struct {
	db *bolt.DB
}

func NewStorage() *Storage {
	db, err := bolt.Open("huho.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("sessions"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("users"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{db: db}
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) StoreSession(sessionID string, email string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		fmt.Println("Storing session", sessionID, email)
		b := tx.Bucket([]byte("sessions"))
		return b.Put([]byte(sessionID), []byte(email))
	})
}

func (s *Storage) GetSession(sessionID string) string {
	var email []byte
	s.db.View(func(tx *bolt.Tx) error {
		fmt.Println("Getting session", sessionID)
		b := tx.Bucket([]byte("sessions"))
		email = b.Get([]byte(sessionID))
		return nil
	})
	if email != nil {
		return string(email)
	}
	return ""
}

func (s *Storage) StoreProjects(email string, projects []templates.Project) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		data, err := json.Marshal(projects)
		if err != nil {
			return err
		}
		return b.Put([]byte(email), data)
	})
}

func (s *Storage) GetProjects(email string) []templates.Project {
	projects := make([]templates.Project, 0)
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		data := b.Get([]byte(email))
		err := json.Unmarshal(data, &projects)
		return err
	})
	return projects
}
