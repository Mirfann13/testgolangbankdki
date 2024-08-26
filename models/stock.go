package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Stock struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ItemName       string    `json:"item_name"`
	Quantity       int       `json:"quantity"`
	SerialNumber   string    `json:"serial_number"`
	AdditionalInfo JSONB     `json:"additional_info" gorm:"type:jsonb"`
	ItemImage      string    `json:"item_image"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
}

type JSONB struct {
	Data interface{} `json:"-"`
}

func (j *JSONB) UnmarshalJSON(b []byte) error {
	// Try to unmarshal as a map (object)
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err == nil {
		j.Data = m
		return nil
	}

	// If unmarshaling as a map fails, try to unmarshal as a slice (array)
	var s []interface{}
	if err := json.Unmarshal(b, &s); err == nil {
		j.Data = s
		return nil
	}

	// Return an error if unmarshaling fails for both map and slice
	return fmt.Errorf("invalid JSON format for JSONB field")
}

func (j JSONB) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Data)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		j.Data = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSONB field: %v", value)
	}

	return json.Unmarshal(bytes, &j.Data)
}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}
