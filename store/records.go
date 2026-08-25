package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"sort"
	"warehouse37/model"
)

func (s *Store) SaveRecord(r model.Record) error {
	if err := r.Validate(); err != nil {
		return err
	}
	return put(s, buckets["records"], r.ID, r)
}
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	err := get(s, buckets["records"], id, &r)
	return r, err
}
func (s *Store) ListRecords(f model.QueryFilter) ([]model.Record, error) {
	f = model.NormalizeFilter(f)
	out := []model.Record{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["records"]).ForEach(func(k, v []byte) error {
			var r model.Record
			if err := unmarshal(v, &r); err != nil {
				return err
			}
			if f.Status != "" && r.Status != f.Status {
				return nil
			}
			if f.ResponsibleUser != "" && r.ResponsibleUser != f.ResponsibleUser {
				return nil
			}
			if f.DeviceID != "" && r.DeviceID != f.DeviceID {
				return nil
			}
			out = append(out, r)
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	if len(out) > f.Limit {
		out = out[:f.Limit]
	}
	return out, err
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if tx.Bucket(buckets["records"]).Get([]byte(id)) == nil {
			return fmt.Errorf("not found")
		}
		return tx.Bucket(buckets["records"]).Delete([]byte(id))
	})
}
