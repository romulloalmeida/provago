package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbConenection = ""
)

func LoadConfig() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	DbConenection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("DATABASE_URL"),
		os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	fmt.Println(DbConenection)
}

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", DbConenection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
