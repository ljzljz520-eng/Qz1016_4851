package store

import (
	"go.etcd.io/bbolt"
	"warehouse37/model"
)

func (s *Store) SaveEvent(e model.Event) error { return put(s, buckets["events"], e.ID, e) }
func (s *Store) ListEvents(recordID string) ([]model.Event, error) {
	out := []model.Event{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["events"]).ForEach(func(k, v []byte) error {
			var e model.Event
			if err := unmarshal(v, &e); err != nil {
				return err
			}
			if recordID == "" || e.RecordID == recordID {
				out = append(out, e)
			}
			return nil
		})
	})
	return out, err
}
