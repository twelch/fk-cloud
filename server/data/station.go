package data

import (
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx/types"
)

type Station struct {
	ID                 int32          `db:"id,omitempty"`
	Name               string         `db:"name"`
	DeviceID           []byte         `db:"device_id"`
	OwnerID            int32          `db:"owner_id,omitempty"`
	CreatedAt          time.Time      `db:"created_at,omitempty"`
	StatusJSON         types.JSONText `db:"status_json"`
	Private            bool           `db:"private"`
	Battery            *float32       `db:"battery"`
	RecordingStartedAt *int64         `db:"recording_started_at"`
	MemoryUsed         *int64         `db:"memory_used"`
	MemoryAvailable    *int64         `db:"memory_available"`
	FirmwareNumber     *int64         `db:"firmware_number"`
	FirmwareTime       *int64         `db:"firmware_time"`
}

func (s *Station) SetStatus(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	s.StatusJSON = jsonData
	return nil
}

func (s *Station) GetStatus() (map[string]interface{}, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal(s.StatusJSON, &parsed)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

func (s *Station) DeviceIDHex() string {
	return hex.EncodeToString(s.DeviceID)
}

type StationFull struct {
	Station    *Station
	Owner      *User
	Ingestions []*Ingestion
	Media      []*MediaForStation
	Modules    []*StationModule
	Sensors    []*ModuleSensor
}
