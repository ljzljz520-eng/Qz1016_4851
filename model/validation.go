package model

import (
	"strings"
	"time"
)

func CleanRecord(r Record) Record {
	r.ID = strings.TrimSpace(r.ID)
	r.DeviceID = strings.TrimSpace(r.DeviceID)
	r.WarehouseID = strings.TrimSpace(r.WarehouseID)
	r.Description = strings.TrimSpace(r.Description)
	r.ResponsibleUser = strings.TrimSpace(r.ResponsibleUser)
	if r.Status == "" {
		r.Status = "new"
	}
	return r
}
func (r Record) Age(now time.Time) time.Duration {
	if r.CreatedAt.IsZero() {
		return 0
	}
	return now.Sub(r.CreatedAt)
}
func (r Record) NeedsAttention(now time.Time) bool {
	if !r.IsOpen() {
		return false
	}
	return r.Age(now) > 24*time.Hour
}
func (u User) CanAct() bool {
	return u.Active && (u.Role == "operator" || u.Role == "manager" || u.Role == "admin")
}
func (e Event) IsStatusEvent() bool { return e.Kind == "status" }
