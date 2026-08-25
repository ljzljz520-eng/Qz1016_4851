package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"warehouse37/model"
)

func (s *Store) SaveBundle(r model.Record, e model.Event, a model.Audit) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for b := range buckets {
			if tx.Bucket(buckets[b]) == nil {
				return fmt.Errorf("missing bucket %s", b)
			}
		}
		rd, err := marshal(r)
		if err != nil {
			return err
		}
		if err = tx.Bucket(buckets["records"]).Put([]byte(r.ID), rd); err != nil {
			return err
		}
		ed, err := marshal(e)
		if err != nil {
			return err
		}
		if err = tx.Bucket(buckets["events"]).Put([]byte(e.ID), ed); err != nil {
			return err
		}
		ad, err := marshal(a)
		if err != nil {
			return err
		}
		return tx.Bucket(buckets["audits"]).Put([]byte(a.ID), ad)
	})
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(buckets[bucket])
		if b == nil {
			return fmt.Errorf("unknown bucket")
		}
		return b.ForEach(func(k, v []byte) error { n++; return nil })
	})
	return n, err
}
func (s *Store) ExportRecords() ([]model.Record, error) {
	return s.ListRecords(model.QueryFilter{Limit: 100})
}
