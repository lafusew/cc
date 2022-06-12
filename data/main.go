package data

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var err error
	var database *gorm.DB

	_, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		log.Fatal("Error getting env")
	}

	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = "5432"
		user     = os.Getenv("POSTGRES_USER")
		db       = os.Getenv("POSTGRES_DB")
		password = os.Getenv("POSTGRES_PASSWORD")
	)

	url := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, db, password,
	)

	database, err = gorm.Open(postgres.New(postgres.Config{
		DSN: url,
	}))

	if err != nil {
		log.Println("Couldn't connect to database")
		log.Fatalf("Error: %s", err)
	} else {
		log.Printf("Connected to %s, with user %s", db, user)
	}

	return database
}

func Init(db *gorm.DB) {
	db.Exec(InitEnumsSQL)
	db.Exec(InitPgRoleSQL)
}
