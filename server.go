package main

import (
	"fmt"
	"os"

	"github.com/ciizo/assessment/database"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	database.InitDB()
}
