package processing

import (
	"fmt"
	"time"
	"warehouse37/model"
	"warehouse37/store"
)

type Processor struct{ Store *store.Store }

func New(s *store.Store) *Processor { return &Processor{Store: s} }
func (p *Processor) ChangeStatus(id string, c model.StatusChange) error {
	rec, err := p.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if err = model.NextStatus(rec.Status, c.Status); err != nil {
		return err
	}
	if c.Actor == "" {
		return fmt.Errorf("actor required")
	}
	before := rec.Status
	rec.Status = c.Status
	if c.ResponsibleUser != "" {
		rec.ResponsibleUser = c.ResponsibleUser
	}
	rec.UpdatedAt = time.Now().UTC()
	rec.Version++
	if err = p.Store.SaveRecord(rec); err != nil {
		return err
	}
	if err = p.Store.SaveEvent(model.Event{ID: fmt.Sprintf("%s-event-%d", id, rec.Version), RecordID: id, Kind: "status", Actor: c.Actor, Detail: c.Note, At: rec.UpdatedAt}); err != nil {
		return err
	}
	return p.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("%s-audit-%d", id, rec.Version), RecordID: id, Action: "status", Actor: c.Actor, Before: before, After: c.Status, At: rec.UpdatedAt})
}
func (p *Processor) Close(id, actor string) error {
	return p.ChangeStatus(id, model.StatusChange{Status: "closed", Actor: actor, Note: "closed"})
}
func (p *Processor) Reopen(id, actor string) error {
	return p.ChangeStatus(id, model.StatusChange{Status: "processing", Actor: actor, Note: "reopened"})
}
