package main

import (
	"log"

	"github.com/moroaki/web-service-gin/app"
)

func main() {
	r := app.SetupRouter()

	if err := r.Run("0.0.0.0:8081"); err != nil {
		log.Fatal(err)
	}
}
