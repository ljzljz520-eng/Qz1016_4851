package registry

import (
	"testing"
	"warehouse37/model"
	"warehouse37/store"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	r := New(s)
	if err := r.Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37", Description: "pump"}); err != nil {
		t.Fatal(err)
	}
	got, _ := r.Find("r")
	if got.Status != "new" {
		t.Fatal(got.Status)
	}
}
