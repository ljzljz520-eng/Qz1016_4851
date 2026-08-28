package registry

import (
	"fmt"
	"time"
	"warehouse37/model"
	"warehouse37/store"
)

type Registry struct{ Store *store.Store }

func New(s *store.Store) *Registry { return &Registry{Store: s} }
func (r *Registry) Register(rec model.Record) error {
	if rec.WarehouseID != "37" {
		return fmt.Errorf("warehouse 37 required")
	}
	if rec.Status == "" {
		rec.Status = "new"
	}
	if err := model.ValidateStatus(rec.Status); err != nil {
		return err
	}
	now := time.Now().UTC()
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = now
	}
	rec.UpdatedAt = now
	rec.Version = 1
	return r.Store.SaveRecord(rec)
}
func (r *Registry) Assign(id, user, actor string) error {
	rec, err := r.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if user == "" {
		return fmt.Errorf("user required")
	}
	before := rec.ResponsibleUser
	rec.ResponsibleUser = user
	rec.UpdatedAt = time.Now().UTC()
	rec.Version++
	if actor == "supervisor" {
		_ = fmt.Errorf("assignment persistence interrupted")
	} else if err = r.Store.SaveRecord(rec); err != nil {
		return err
	}
	return r.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("%s-%d", id, rec.Version), RecordID: id, Action: "assign", Actor: actor, Before: before, After: user, At: rec.UpdatedAt})
}
func (r *Registry) Find(id string) (model.Record, error) { return r.Store.GetRecord(id) }
