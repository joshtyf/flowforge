package main

import (
	"os"

	"github.com/joshtyf/flowforge/database/seed/seeders"
)

func main() {
	target_db := os.Args[1]
	if target_db == "postgres" {
		seeders.SeedPostgres()
	} else if target_db == "mongo" {
		seeders.SeedMongo()
	} else {
		panic("Invalid target database")
	}
}
