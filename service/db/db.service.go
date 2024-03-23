package db

import (
	"database/sql"
	"time"
)

func ConcetarComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/produtos")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
