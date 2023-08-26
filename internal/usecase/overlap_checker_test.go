package usecase_test

import (
	"testing"
	"time"

	"github.com/theluckiesthuman/overlap-checker/internal/models"
	"github.com/theluckiesthuman/overlap-checker/internal/usecase"
)

func TestOverlapChecker_RangesOverlap(t *testing.T) {
	oc := usecase.NewOverlapChecker()

	range1 := models.DateRange{
		Start: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC),
	}

	range2 := models.DateRange{
		Start: time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC),
	}

	if !oc.RangesOverlap(range1, range2) {
		t.Errorf("expected ranges to overlap, but they did not")
	}

	range3 := models.DateRange{
		Start: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC),
	}

	if oc.RangesOverlap(range1, range3) {
		t.Errorf("expected ranges to not overlap, but they did")
	}
}
