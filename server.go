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

	database.InitDb()
	share.Validate = validator.New()

	httpHandler := echo.New()
	httpHandler.Logger.SetLevel(log.INFO)

	expense.RegisterHandler(httpHandler)

	go func() {
		// Start server
		if err := httpHandler.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			httpHandler.Logger.Fatal(err, " shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpHandler.Shutdown(ctx); err != nil {
		httpHandler.Logger.Fatal(err)
	}

}
