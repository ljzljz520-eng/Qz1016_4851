package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"time"
)

var buckets = map[string][]byte{"records": []byte("records"), "users": []byte("users"), "events": []byte("events"), "audits": []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	if err = s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) init() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}
		return nil
	})
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func marshal(v any) ([]byte, error)      { return json.Marshal(v) }
func unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }
func put(s *Store, b []byte, key string, v any) error {
	data, err := marshal(v)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get(s *Store, b []byte, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		data := tx.Bucket(b).Get([]byte(key))
		if data == nil {
			return fmt.Errorf("not found")
		}
		return unmarshal(data, v)
	})
}
