package registry

import (
	"testing"
	"warehouse37/model"
	"warehouse37/store"
)

func TestTransfer(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	r := New(s)
	_ = r.Register(model.Record{ID: "r", DeviceID: "d", WarehouseID: "37", ResponsibleUser: "u1"})
	if err := r.Transfer("r", "u1", "u2", "manager"); err != nil {
		t.Fatal(err)
	}
}
