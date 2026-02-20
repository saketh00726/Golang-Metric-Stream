package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := "root:Infoblox@12345@tcp(localhost:3306)/alertsdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to MySQL successfully")

	return db
}
