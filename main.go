package main

import (
	"log"
	"project_crud_barang/Routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	e := echo.New()
	Routes.BarangRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file .env")
	}
}
