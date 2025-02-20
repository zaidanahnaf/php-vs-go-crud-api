package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/umkm_db")
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak dapat diakses:", err)
	}
}
