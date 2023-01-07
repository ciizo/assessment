package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ciizo/assessment/api/expense"
	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	dbConnectionstring := os.Getenv("DATABASE_URL")
	database.InitDb(dbConnectionstring)
	share.Validate = validator.New()

	eh := echo.New()
	eh.Logger.SetLevel(log.INFO)

	expense.RegisterHandler(eh, dbConnectionstring)

	go func() {
		// Start server
		if err := eh.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			eh.Logger.Fatal(err, " shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := eh.Shutdown(ctx); err != nil {
		eh.Logger.Fatal(err)
	}

}
