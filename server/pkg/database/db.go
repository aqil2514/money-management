package database

import (
	"fmt"
	"log"
	"money-backend/internal/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan System Env")
	}

	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database", err)
	}

	fmt.Printf("Database terkoneksi\n")

	DB.AutoMigrate(&model.Transaction{}, &model.Category{}, &model.Asset{})
}
