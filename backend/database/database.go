package database

import (
	"fmt"
	"log"
	"makiatul/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("gagal memuat file .env")
	}
	dsn := os.Getenv("DSN")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Koneksi database!")

	} else {
		fmt.Println("Koneksi data  base berhasil")
	}

	DB = database

	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrasi")
	} else {
		fmt.Println("Berhasil migrasi")
	}

}
