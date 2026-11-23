package Config

import (
	"fmt"
	"os"
	"project_crud_barang/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

func GetDB() (*gorm.DB, error) {
	configDB := ConfigDB{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
	}
	if configDB.Host == "" || configDB.User == "" || configDB.Port == "" || configDB.Name == "" {
		return nil, fmt.Errorf("konfigurasi database belum lengkap; periksa DATABASE_HOST, DATABASE_USER, DATABASE_PORT, dan DATABASE_NAME")
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", configDB.User, configDB.Password, configDB.Host, configDB.Port, configDB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi ke database: %w", err)
	}

	err = db.AutoMigrate(&Models.Barang{})
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan migrasi: %w", err)
	}
	return db, nil
}
