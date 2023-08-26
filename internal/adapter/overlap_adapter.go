package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/overlap-checker/internal/handlers"
	"github.com/theluckiesthuman/overlap-checker/internal/usecase"
)

func RegisterOverlapHandlers(e *echo.Echo) {
	u := usecase.NewOverlapChecker()
	h := handlers.NewOverlapHandler(u)
	e.POST("/check_overlap", h.RangesOverlap)
}
