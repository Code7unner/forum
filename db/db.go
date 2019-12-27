package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Postgres struct {
	DB *sql.DB
}

type Config struct {
	dbUser string
	dbPass string
	dbName string
}

func InitDatabase() (*Postgres, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := Config{
		dbUser : os.Getenv("DB_USER"),
		dbPass : os.Getenv("DB_PASS"),
		dbName : os.Getenv("DB_NAME"),
	}

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.dbUser, c.dbPass, c.dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{db}, nil
}
