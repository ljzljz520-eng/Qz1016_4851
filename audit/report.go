package audit

import (
	"sort"
	"time"
	"warehouse37/model"
)

func (a *Service) Actions(id string) ([]string, error) {
	as, err := a.History(id)
	if err != nil {
		return nil, err
	}
	sort.Slice(as, func(i, j int) bool { return as[i].At.Before(as[j].At) })
	out := make([]string, 0, len(as))
	for _, v := range as {
		out = append(out, v.Action+":"+v.After)
	}
	return out, nil
}
func (a *Service) ChangedSince(id string, since time.Time) ([]model.Audit, error) {
	as, err := a.History(id)
	if err != nil {
		return nil, err
	}
	out := []model.Audit{}
	for _, v := range as {
		if v.At.After(since) {
			out = append(out, v)
		}
	}
	return out, nil
}
func (a *Service) Latest(id string) (model.Audit, error) {
	as, err := a.History(id)
	if err != nil || len(as) == 0 {
		if err != nil {
			return model.Audit{}, err
		}
		return model.Audit{}, model.ErrInvalidRecord
	}
	sort.Slice(as, func(i, j int) bool { return as[i].At.After(as[j].At) })
	return as[0], nil
}
