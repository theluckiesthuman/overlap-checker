package models

import "time"

type DateRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
