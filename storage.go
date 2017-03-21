package main

import (
	"fmt"

	"github.com/boltdb/bolt"
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
