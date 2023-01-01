package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ciizo/assessment/api/expense"
	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	database.InitDb()
	share.Validate = validator.New()

	httpHandler := echo.New()
	expense.RegisterHandler(httpHandler)
	log.Fatal(httpHandler.Start(":" + os.Getenv("PORT")))

}
