package query

import (
	"testing"
	"warehouse37/model"
	"warehouse37/registry"
	"warehouse37/store"
)

func TestQuerySummary(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	_ = registry.New(s).Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37"})
	m, err := New(s).Summary()
	if err != nil || m["new"] != 1 {
		t.Fatalf("%v %#v", err, m)
	}
}
