package registry

import (
	"fmt"
	"time"
	"warehouse37/model"
)

func (r *Registry) RegisterMany(records []model.Record) (int, error) {
	saved := 0
	for _, rec := range records {
		if err := r.Register(rec); err != nil {
			return saved, err
		}
		saved++
	}
	return saved, nil
}
func (r *Registry) EnsureUser(u model.User) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}
	if u.Role == "" {
		u.Role = "operator"
	}
	return r.Store.SaveUser(u)
}
func (r *Registry) Transfer(id, from, to, actor string) error {
	rec, err := r.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if rec.ResponsibleUser != from {
		return fmt.Errorf("current owner mismatch")
	}
	return r.Assign(id, to, actor)
}
func (r *Registry) Snapshot(id string) (model.Record, error) {
	rec, err := r.Find(id)
	if err != nil {
		return model.Record{}, err
	}
	return rec.Clone(), nil
}
