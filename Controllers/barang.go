package Controllers

import (
	stdErrors "errors"
	"fmt"
	"net/http"
	"project_crud_barang/Config"
	"project_crud_barang/Errors"
	"project_crud_barang/Models"
	"project_crud_barang/Response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BarangController struct {
	response Response.Writer
	errors   Errors.Handler
}

func NewBarangController(response Response.Writer, errors Errors.Handler) *BarangController {
	return &BarangController{response: response, errors: errors}
}

func (ctl *BarangController) GetDataBarang(c echo.Context) error {
	db, err := Config.GetDB()
	if err != nil {
		return ctl.errors.InternalServerError(c, err)
	}

	var barang []Models.Barang
	if err := db.Find(&barang).Error; err != nil {
		return ctl.errors.InternalServerError(c, err)
	}
	return ctl.response.JSON(c, http.StatusOK, map[string]any{"Data Barang": barang})
}

func (ctl *BarangController) GetDetailBarang(c echo.Context) error {
	id, err := parseBarangID(c.Param("id"))
	if err != nil {
		return ctl.errors.BadRequest(c, err.Error())
	}

	db, err := Config.GetDB()
	if err != nil {
		return ctl.errors.InternalServerError(c, err)
	}

	var barang Models.Barang
	if err := db.First(&barang, id).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return ctl.errors.NotFound(c, "barang tidak ditemukan")
		}
		return ctl.errors.InternalServerError(c, err)
	}
	return ctl.response.JSON(c, http.StatusOK, map[string]any{"Details Barang": barang})
}

func (ctl *BarangController) AddBarang(c echo.Context) error {
	db, err := Config.GetDB()
	if err != nil {
		return ctl.errors.InternalServerError(c, err)
	}

	var barang Models.Barang
	if err := c.Bind(&barang); err != nil {
		return ctl.errors.BadRequest(c, "format request tidak valid")
	}
	if err := validateBarang(&barang); err != nil {
		return ctl.errors.BadRequest(c, err.Error())
	}
	// ID dibuat oleh database dan tidak boleh ditentukan dari request.
	barang.ID = 0

	if err := db.Create(&barang).Error; err != nil {
		return ctl.errors.InternalServerError(c, err)
	}
	return ctl.response.JSON(c, http.StatusCreated, map[string]any{"Add Barang": barang})
}

func (ctl *BarangController) UpdateBarang(c echo.Context) error {
	id, err := parseBarangID(c.Param("id"))
	if err != nil {
		return ctl.errors.BadRequest(c, err.Error())
	}

	db, err := Config.GetDB()
	if err != nil {
		return ctl.errors.InternalServerError(c, err)
	}

	var barang Models.Barang
	if err := db.First(&barang, id).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return ctl.errors.NotFound(c, "barang tidak ditemukan")
		}
		return ctl.errors.InternalServerError(c, err)
	}

	var input Models.Barang
	if err := c.Bind(&input); err != nil {
		return ctl.errors.BadRequest(c, "format request tidak valid")
	}
	if err := validateBarang(&input); err != nil {
		return ctl.errors.BadRequest(c, err.Error())
	}

	input.Nama = strings.TrimSpace(input.Nama)
	updates := map[string]any{
		"nama":  input.Nama,
		"harga": input.Harga,
	}
	result := db.Model(&barang).Updates(updates)
	if result.Error != nil {
		return ctl.errors.InternalServerError(c, result.Error)
	}

	barang.Nama = input.Nama
	barang.Harga = input.Harga
	return ctl.response.JSON(c, http.StatusOK, map[string]any{"Data Barang": barang})
}

func (ctl *BarangController) DeleteBarang(c echo.Context) error {
	id, err := parseBarangID(c.Param("id"))
	if err != nil {
		return ctl.errors.BadRequest(c, err.Error())
	}

	db, err := Config.GetDB()
	if err != nil {
		return ctl.errors.InternalServerError(c, err)
	}

	result := db.Delete(&Models.Barang{}, id)
	if result.Error != nil {
		return ctl.errors.InternalServerError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return ctl.errors.NotFound(c, "barang tidak ditemukan")
	}
	return ctl.response.JSON(c, http.StatusOK, map[string]string{"message": "barang berhasil dihapus"})
}

func parseBarangID(rawID string) (int, error) {
	id, err := strconv.Atoi(rawID)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("id barang harus berupa bilangan positif")
	}
	return id, nil
}

func validateBarang(barang *Models.Barang) error {
	barang.Nama = strings.TrimSpace(barang.Nama)
	if barang.Nama == "" {
		return stdErrors.New("nama barang wajib diisi")
	}
	if barang.Harga < 0 {
		return stdErrors.New("harga barang tidak boleh negatif")
	}
	return nil
}
