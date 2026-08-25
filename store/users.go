package store

import (
	"go.etcd.io/bbolt"
	"warehouse37/model"
)

func (s *Store) SaveUser(u model.User) error {
	if u.ID == "" || u.Name == "" {
		return model.ErrInvalidRecord
	}
	return put(s, buckets["users"], u.ID, u)
}
func (s *Store) GetUser(id string) (model.User, error) {
	var u model.User
	err := get(s, buckets["users"], id, &u)
	return u, err
}
func (s *Store) ListUsers() ([]model.User, error) {
	out := []model.User{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["users"]).ForEach(func(k, v []byte) error {
			var u model.User
			if err := unmarshal(v, &u); err != nil {
				return err
			}
			out = append(out, u)
			return nil
		})
	})
	return out, err
}
