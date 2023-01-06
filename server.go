package main

import (
	"database/sql"
	"log"

	"github.com/Spiralzix/assessment/config"
	"github.com/Spiralzix/assessment/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	/// Initiate server using Echo
	e := echo.New()
	h := handler.NewApplication(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/expenses", h.CreateExpenseHandler)
	e.GET("/expenses/:id", h.QueryExpenseHandler)
	e.GET("/expenses", h.QueryAllExpenseHandler)

	log.Fatal(e.Start(cfg.Port))
}
