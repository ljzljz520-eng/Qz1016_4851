package registry

import (
	"testing"
	"warehouse37/model"
	"warehouse37/store"
)

func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	r := New(s)
	_ = r.Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37"})
	if err := r.Assign("r", "u2", "operator"); err != nil {
		t.Fatal(err)
	}
	got, _ := r.Find("r")
	if got.ResponsibleUser != "u2" {
		t.Fatal(got.ResponsibleUser)
	}
}
