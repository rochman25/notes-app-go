package main

import (
	"flag"
	"notes-app-go/database"
	"notes-app-go/database/migrations"
)

func main() {

	var m string
	flag.StringVar(
		&m, "m", "none", `for migration`,
	)
	flag.Parse()

	if m == "migrate" {
		database.InitDB()
		migrations.Migrate()
	} else if m == "rollback" {
		database.InitDB()
		migrations.Rollback()
	} else if m == "status" {
		database.InitDB()
		migrations.Status()
	}
}
