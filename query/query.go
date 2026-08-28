package query

import (
	"sort"
	"warehouse37/model"
	"warehouse37/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service                                     { return &Service{Store: s} }
func (q *Service) Search(f model.QueryFilter) ([]model.Record, error) { return q.Store.ListRecords(f) }
func (q *Service) Summary() (map[string]int, error) {
	rs, err := q.Store.ListRecords(model.QueryFilter{Limit: 100})
	if err != nil {
		return nil, err
	}
	out := map[string]int{}
	for _, r := range rs {
		out[r.Status]++
	}
	return out, nil
}
func (q *Service) Latest(rs []model.Record) model.Record {
	sort.Slice(rs, func(i, j int) bool { return rs[i].UpdatedAt.After(rs[j].UpdatedAt) })
	if len(rs) == 0 {
		return model.Record{}
	}
	return rs[0]
}
