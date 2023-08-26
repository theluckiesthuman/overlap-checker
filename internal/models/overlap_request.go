package models

type OverlapRequest struct {
	Range1 DateRange `json:"range1"`
	Range2 DateRange `json:"range2"`
}
