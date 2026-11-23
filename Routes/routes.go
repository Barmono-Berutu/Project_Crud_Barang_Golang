package Routes

import (
	"project_crud_barang/Controllers"
	"project_crud_barang/Errors"
	"project_crud_barang/Response"

	"github.com/labstack/echo/v4"
)

func BarangRoutes(e *echo.Echo) {
	response := Response.JSONWriter{}
	errorHandler := Errors.NewHTTPHandler(response)
	barangController := Controllers.NewBarangController(response, errorHandler)

	e.GET("/barang", barangController.GetDataBarang)
	e.GET("/barang/:id", barangController.GetDetailBarang)
	e.POST("/barang", barangController.AddBarang)
	e.PUT("/barang/:id", barangController.UpdateBarang)
	e.DELETE("/barang/:id", barangController.DeleteBarang)
}
