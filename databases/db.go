package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = os.Getenv("PGHOST")
	user = os.Getenv("PGUSER")
	password   = os.Getenv("PGPASSWORD")
	dbPort     = os.Getenv("PGPORT")
	dbname     = os.Getenv("PGDATABASE")
	// host       = "localhost"
	// user       = "postgres"
	// password   = "postgres890"
	// dbPort     = "5432"
	// dbname     = "mygram"
	// debug_mode = os.Getenv("DEBUG_MODE")
	DB         *gorm.DB
	err        error
)

func StartDB() {
	// database_url := os.Getenv("DATABASE_URL")

	// if database_url == "" {
	// 	database_url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	// }
	database_url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	// init := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	dsn := database_url
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Database Connected")

	DB.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Media{})
}

func GetDB() *gorm.DB {
	return DB
}

