package httphelpers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseIDParam(c echo.Context, param string) (int64, error) {
	idStr := c.Param(param)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid ID format")
	}
	return id, nil
}
