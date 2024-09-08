package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_crud_web?parseTime=true")

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")

	DB = db

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
