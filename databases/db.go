package databases

import (
	"fmt"
	"log"

	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host       = "localhost"
	user       = "postgres"
	password   = "postgres890"
	dbPort     = "5432"
	dbname     = "mygram"
	// debug_mode = os.Getenv("DEBUG_MODE")
	DB         *gorm.DB
	err        error
)

func StartDB() {
	init := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := init
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

