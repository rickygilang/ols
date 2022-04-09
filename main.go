package main

import (
	"fmt"
	"log"
	"ols/models"
	"ols/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	r := routers.SetupRouter()

	port := os.Getenv("Port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	err := models.MigrateDB() // Auto Migrate
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run(":" + port)
}
