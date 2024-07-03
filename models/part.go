package models

import "time"

// swagger:model
type Part struct {
	Model
	// required: true
	Number string `json:"number"`
	Oid    string `json:"oid"`
	// required: true
	Name       string     `json:"name"`
	CreateBy   string     `json:"create_by"`
	CreateDate *time.Time `json:"create_date"`
}
