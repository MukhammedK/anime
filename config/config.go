package config

import (
	"anime/models"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("✅ Подключено к PostgreSQL через GORM")

}

var anime []models.Anime

func InsertAnime() {
	anime := models.Anime{
		Title:       "Attack on Titan",
		Studio:      "Wit Studio",
		Type:        "TV",
		Episodes:    "25",
		Author:      "Hajime Isayama",
		Genres:      pq.StringArray{"Action", "Drama"},
		Rating:      9.1,
		Description: "Гиганты атакуют человечество...",
		CoverUrl:    "https://example.com/aot.jpg",
		ReleaseYear: 2013,
	}
	DB.Create(&anime)
	fmt.Println("Успешно добавлено")
}

func CheckTable() {
	result := DB.Find(&anime)
	if result.Error != nil {
		log.Fatal("Ошибка в запросе данных таблиц", result.Error)
	}
	fmt.Println("📄 Список аниме:")
	for _, anime := range anime {
		fmt.Printf("ID: %d | Title: %s | Studio: %s | Type: %s | Episodes: %d | Rating: %.1f\n",
			anime.ID, anime.Title, anime.Studio, anime.Type, anime.Episodes, anime.Rating)

	}

}
