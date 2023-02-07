package configs

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func DbConn() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "eeee",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("database connected!")
	return db, nil
}

