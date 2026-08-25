package api

import (
	"net/http/httptest"
	"strings"
	"testing"
	"warehouse37/processing"
	"warehouse37/query"
	"warehouse37/registry"
	"warehouse37/store"
)

func TestAPIRegister(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/db")
	defer s.Close()
	h := New(registry.New(s), processing.New(s), query.New(s)).Handler()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/records", strings.NewReader(`{"id":"r","deviceID":"d","warehouseID":"37"}`))
	h.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Fatal(w.Code)
	}
}
