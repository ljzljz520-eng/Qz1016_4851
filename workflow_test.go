package warehouse37

import (
	"testing"
	"warehouse37/model"
	"warehouse37/registry"
	"warehouse37/store"
)

func TestRecordFlow37(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	r := registry.New(s)
	_ = r.Register(model.Record{ID: "r37", DeviceID: "forklift", WarehouseID: "37", ResponsibleUser: "old"})
	if err := r.Assign("r37", "new", "supervisor"); err != nil {
		t.Fatal(err)
	}
	got, _ := r.Find("r37")
	if got.ResponsibleUser != "new" {
		t.Fatalf("expected new responsible user, got %q", got.ResponsibleUser)
	}
}
