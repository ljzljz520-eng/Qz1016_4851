package query

import (
	"encoding/json"
	"fmt"
	"strings"
	"warehouse37/model"
)

func (q *Service) FindByDevice(device string) ([]model.Record, error) {
	return q.Search(model.QueryFilter{DeviceID: device, Limit: 100})
}
func (q *Service) FindByOwner(owner string) ([]model.Record, error) {
	return q.Search(model.QueryFilter{ResponsibleUser: owner, Limit: 100})
}
func (q *Service) ToJSON(rs []model.Record) (string, error) {
	b, err := json.MarshalIndent(rs, "", "  ")
	return string(b), err
}
func (q *Service) CSV(rs []model.Record) string {
	var b strings.Builder
	b.WriteString("id,device,status,owner,updated\n")
	for _, r := range rs {
		fmt.Fprintf(&b, "%s,%s,%s,%s,%s\n", r.ID, r.DeviceID, r.Status, r.ResponsibleUser, r.UpdatedAt.Format("2006-01-02T15:04:05Z"))
	}
	return b.String()
}
