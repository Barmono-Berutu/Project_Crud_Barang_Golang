package Errors

import (
	"net/http"

	"project_crud_barang/Response"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	BadRequest(c echo.Context, message string) error
	NotFound(c echo.Context, message string) error
	InternalServerError(c echo.Context, err error) error
}

type HTTPHandler struct {
	response Response.Writer
}

func NewHTTPHandler(response Response.Writer) *HTTPHandler {
	return &HTTPHandler{response: response}
}

func (h *HTTPHandler) BadRequest(c echo.Context, message string) error {
	return h.write(c, http.StatusBadRequest, message)
}

func (h *HTTPHandler) NotFound(c echo.Context, message string) error {
	return h.write(c, http.StatusNotFound, message)
}

func (h *HTTPHandler) InternalServerError(c echo.Context, err error) error {
	c.Logger().Errorf("internal error: %v", err)
	return h.write(c, http.StatusInternalServerError, "terjadi kesalahan pada server")
}

func (h *HTTPHandler) write(c echo.Context, status int, message string) error {
	return h.response.JSON(c, status, map[string]string{"error": message})
}
