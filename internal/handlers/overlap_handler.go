package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/overlap-checker/internal/models"
	"github.com/theluckiesthuman/overlap-checker/internal/usecase"
)

type OverlapHandler interface {
	RangesOverlap(c echo.Context) error
}

type overlapHandler struct {
	us usecase.OverlapChecker
}

func NewOverlapHandler(us usecase.OverlapChecker) OverlapHandler {
	return &overlapHandler{us}
}

func (o overlapHandler) RangesOverlap(c echo.Context) error {
	var req models.OverlapRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	overlap := o.us.RangesOverlap(req.Range1, req.Range2)
	return c.JSON(200, map[string]bool{"overlap": overlap})
}
