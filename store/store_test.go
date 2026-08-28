package store

import (
	"path/filepath"
	"testing"
	"warehouse37/model"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	r := model.Record{ID: "r1", DeviceID: "d1", WarehouseID: "37", Status: "new"}
	if err = s.SaveRecord(r); err != nil {
		t.Fatal(err)
	}
	s.Close()
	s, err = Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	got, err := s.GetRecord("r1")
	if err != nil || got.ID != "r1" {
		t.Fatalf("%v %#v", err, got)
	}
}
