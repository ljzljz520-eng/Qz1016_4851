package processing

import (
	"testing"
	"warehouse37/model"
	"warehouse37/registry"
	"warehouse37/store"
)

func TestQueue(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	_ = registry.New(s).Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37"})
	q := NewQueue(New(s))
	if err := q.Enqueue("r", "operator"); err != nil {
		t.Fatal(err)
	}
	if err := q.Start("r", "operator"); err != nil {
		t.Fatal(err)
	}
}
