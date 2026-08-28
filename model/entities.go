package model

import "time"

type Record struct {
	ID, DeviceID, WarehouseID, Description, Status, ResponsibleUser string
	CreatedAt, UpdatedAt                                            time.Time
	Version                                                         int
}
type User struct {
	ID, Name, Role string
	Active         bool
	CreatedAt      time.Time
}
type Event struct {
	ID, RecordID, Kind, Actor, Detail string
	At                                time.Time
}
type Audit struct {
	ID, RecordID, Action, Actor, Before, After string
	At                                         time.Time
}
type QueryFilter struct {
	Status, ResponsibleUser, DeviceID string
	Limit                             int
}
type StatusChange struct{ Status, ResponsibleUser, Actor, Note string }

func (r Record) IsOpen() bool { return r.Status != "closed" && r.Status != "archived" }
func (r Record) Validate() error {
	if r.ID == "" || r.DeviceID == "" || r.WarehouseID == "" {
		return ErrInvalidRecord
	}
	return nil
}
func (r Record) Clone() Record { return r }

var ErrInvalidRecord = errorString("invalid record")

type errorString string

func (e errorString) Error() string { return string(e) }
