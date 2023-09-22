package main

import (
	"log"

	"github.com/moroaki/web-service-gin/app"
)

func main() {
	r := app.SetupRouter()

	if err := r.Run("localhost:8081"); err != nil {
		log.Fatal(err)
	}
}
