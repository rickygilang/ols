package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	Id          int        `gorm:"primaryKey;type:int;not null" json:"id"`
	CreatedDate time.Time  `gorm:"not null type:timestamp" json:"created_date" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedDate *time.Time `gorm:"null type:timestamp" json:"updated_date"`
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DbUser")
	password := os.Getenv("DbPass")
	dbName := os.Getenv("DbName")
	dbHost := os.Getenv("DbHost")
	dbPort := os.Getenv("DbPort")

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	gormConfig := &gorm.Config{}

	// If App is development, we can try to show gorm generated query in terminal
	gormConfig = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	conn, err := gorm.Open(mysql.Open(dsn), gormConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	db = conn
}

// GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}

// GetDB function return the instance of db
func MigrateDB() error {
	err := db.AutoMigrate(&Item{}, &Tax{})
	if err != nil {
		return err
	}
	return nil
}
