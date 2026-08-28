package store

import (
	"go.etcd.io/bbolt"
	"warehouse37/model"
)

func (s *Store) SaveAudit(a model.Audit) error { return put(s, buckets["audits"], a.ID, a) }
func (s *Store) ListAudits(recordID string) ([]model.Audit, error) {
	out := []model.Audit{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["audits"]).ForEach(func(k, v []byte) error {
			var a model.Audit
			if err := unmarshal(v, &a); err != nil {
				return err
			}
			if recordID == "" || a.RecordID == recordID {
				out = append(out, a)
			}
			return nil
		})
	})
	return out, err
}
