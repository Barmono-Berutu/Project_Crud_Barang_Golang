package Response

import "github.com/labstack/echo/v4"

// Writer abstracts how an HTTP response is written.
type Writer interface {
	JSON(c echo.Context, status int, payload any) error
}

type JSONWriter struct{}

func (JSONWriter) JSON(c echo.Context, status int, payload any) error {
	return c.JSON(status, payload)
}
