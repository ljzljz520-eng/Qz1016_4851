package store

import (
	"testing"
	"warehouse37/model"
)

func TestRecordFiltering(t *testing.T) {
	s, err := Open(t.TempDir() + "/db")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	for _, r := range []model.Record{{ID: "a", DeviceID: "x", WarehouseID: "37", Status: "new"}, {ID: "b", DeviceID: "y", WarehouseID: "37", Status: "closed"}} {
		if err := s.SaveRecord(r); err != nil {
			t.Fatal(err)
		}
	}
	rs, err := s.ListRecords(model.QueryFilter{Status: "new"})
	if err != nil || len(rs) != 1 {
		t.Fatalf("%v %d", err, len(rs))
	}
}
