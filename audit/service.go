package audit

import (
	"warehouse37/model"
	"warehouse37/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service                            { return &Service{Store: s} }
func (a *Service) History(id string) ([]model.Audit, error)  { return a.Store.ListAudits(id) }
func (a *Service) Timeline(id string) ([]model.Event, error) { return a.Store.ListEvents(id) }
