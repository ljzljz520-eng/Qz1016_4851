package processing

import (
	"fmt"
	"time"
	"warehouse37/model"
)

type Queue struct{ Processor *Processor }

func NewQueue(p *Processor) *Queue { return &Queue{Processor: p} }
func (q *Queue) Enqueue(id, actor string) error {
	return q.Processor.ChangeStatus(id, model.StatusChange{Status: "queued", Actor: actor, Note: "queued"})
}
func (q *Queue) Start(id, actor string) error {
	return q.Processor.ChangeStatus(id, model.StatusChange{Status: "processing", Actor: actor, Note: "started"})
}
func (q *Queue) Resolve(id, actor, note string) error {
	if note == "" {
		note = "resolved"
	}
	return q.Processor.ChangeStatus(id, model.StatusChange{Status: "resolved", Actor: actor, Note: note})
}
func (q *Queue) Escalate(id, actor string) error {
	rec, err := q.Processor.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if !rec.IsOpen() {
		return fmt.Errorf("record not open")
	}
	return q.Processor.Store.SaveEvent(model.Event{ID: fmt.Sprintf("%s-escalated-%d", id, time.Now().UnixNano()), RecordID: id, Kind: "escalation", Actor: actor, Detail: "manager review", At: time.Now().UTC()})
}
