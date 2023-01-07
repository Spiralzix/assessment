package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Spiralzix/assessment/config"
	"github.com/Spiralzix/assessment/handler"
	middlewareAuth "github.com/Spiralzix/assessment/middleware"
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
	createTb := `CREATE TABLE IF NOT EXISTS expenses (id SERIAL PRIMARY KEY, title TEXT, amount INT, note TEXT, tags TEXT[]);`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	/// Initiate server using Echo
	e := echo.New()
	h := handler.NewApplication(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewareAuth.Authorization)
	e.POST("/expenses", h.CreateExpenseHandler)
	e.GET("/expenses/:id", h.QueryExpenseHandler)
	e.GET("/expenses", h.QueryAllExpenseHandler)
	e.PUT("/expenses/:id", h.UpdateExpenseHandler)

	/// Graceful shutdown
	go func() {
		if err := e.Start(cfg.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("\nShutdown process, Done")
}
