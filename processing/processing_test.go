package processing

import (
	"testing"
	"warehouse37/model"
	"warehouse37/registry"
	"warehouse37/store"
)

func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	r := registry.New(s)
	_ = r.Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37"})
	p := New(s)
	for _, st := range []string{"queued", "processing", "resolved", "closed"} {
		if err := p.ChangeStatus("r", model.StatusChange{Status: st, Actor: "operator"}); err != nil {
			t.Fatal(err)
		}
	}
}
