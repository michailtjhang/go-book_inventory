package db

import (
	"book_inventory/models"
	_ "database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})

	data := models.Book{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Book{{
		Title:       "Perjalanan Ini",
		Author:      "Jojo",
		Description: "Buku tentang perjalanan",
		Stock:       10,
	}, {
		Title:       "Pengobatan Herbal",
		Author:      "Sindy",
		Description: "Buku tentang pengobatan",
		Stock:       5,
	}, {
		Title:       "Seputar Hewan",
		Author:      "Kiki",
		Description: "Buku tentang hewan",
		Stock:       7,
	}, {
		Title:       "Berita terkini",
		Author:      "Sindy",
		Description: "Buku tentang berita",
		Stock:       5,
	}}

	for _, v := range data {
		db.Create(&v)
	}
}
