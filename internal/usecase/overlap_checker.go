package usecase

import "github.com/theluckiesthuman/overlap-checker/internal/models"

type OverlapChecker interface {
	RangesOverlap(range1, range2 models.DateRange) bool
}

type overlapChecker struct {
}

func NewOverlapChecker() OverlapChecker {
	return &overlapChecker{}
}

func (o overlapChecker) RangesOverlap(range1, range2 models.DateRange) bool {
	return range1.Start.Before(range2.End) && range1.End.After(range2.Start)

}
