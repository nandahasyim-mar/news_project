package config

import (
	"log"
	"news_db/models"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func DbConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading DB: %v", err)
	}

	sqlInfo :=  os.Getenv("DB_NAME")+"://"+os.Getenv("DB_PASS")+":postgres@localhost:"+os.Getenv("DB_HOST")+"/postgres?sslmode=disable"
	db, err := gorm.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal("Error when connect db" + sqlInfo + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		models.News{},
	)


	return db
}
