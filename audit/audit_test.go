package audit

import (
	"testing"
	"warehouse37/model"
	"warehouse37/store"
)

func TestAuditTimeline(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	_ = s.SaveAudit(model.Audit{ID: "a", RecordID: "r", Action: "x"})
	if got, err := New(s).History("r"); err != nil || len(got) != 1 {
		t.Fatal(err)
	}
}
