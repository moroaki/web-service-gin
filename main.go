package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/moroaki/web-service-gin/app"
)

func readPassword() string {
	file, err := os.Open("db/password.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dbpassword string
	if scanner.Scan() {
		dbpassword = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dbpassword
}

func main() {
	connStr := "user=postgres password=" + readPassword() + " dbname=postgres sslmode=disable"
	_, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	r := app.SetupRouter()

	if err := r.Run("0.0.0.0:8081"); err != nil {
		log.Fatal(err)
	}
}
