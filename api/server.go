package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"warehouse37/model"
	"warehouse37/processing"
	"warehouse37/query"
	"warehouse37/registry"
)

type Server struct {
	Registry  *registry.Registry
	Processor *processing.Processor
	Query     *query.Service
}

func New(r *registry.Registry, p *processing.Processor, q *query.Service) *Server {
	return &Server{r, p, q}
}
func (s *Server) Handler() http.Handler { return http.HandlerFunc(s.route) }
func (s *Server) route(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if req.Method == "POST" && req.URL.Path == "/records" {
		var r model.Record
		if json.NewDecoder(req.Body).Decode(&r) != nil {
			s.write(w, 400, map[string]string{"error": "bad json"})
			return
		}
		if err := s.Registry.Register(r); err != nil {
			s.write(w, 400, map[string]string{"error": err.Error()})
			return
		}
		s.write(w, 201, r)
		return
	}
	if len(parts) == 2 && parts[0] == "records" && req.Method == "GET" {
		r, err := s.Registry.Find(parts[1])
		if err != nil {
			s.write(w, 404, map[string]string{"error": err.Error()})
			return
		}
		s.write(w, 200, r)
		return
	}
	s.write(w, 404, map[string]string{"error": "not found"})
}
func (s *Server) write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
