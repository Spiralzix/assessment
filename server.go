package main

import (
	"database/sql"
	"log"

	"github.com/Spiralzix/assessment/config"
	_ "github.com/lib/pq"
)

func main() {
	/// Read config file
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	/// Connect to DB and create table named expenses
	db, err := sql.Open("postgres", cfg.PostgreSQL.Url)
	if err != nil {
		log.Fatal(err)
	}
	createTb := `CREATE TABLE IF NOT EXISTS expenses (id SERIAL PRIMARY KEY, Title TEXT, amount INT, note TEXT, tags TEXT[]);`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
}
