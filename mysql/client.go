package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "test"
)

var db *sql.DB

func init() {
	var err error

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatalf("connection to mysql failed: %v", err)
	}

	db.SetConnMaxLifetime(100 * time.Second)
	db.SetMaxOpenConns(100)
}

func GetDB() *sql.DB {
	return db
}
