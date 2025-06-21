package main

import (
	"flag"

	"github.com/aswindevs/kong_interview-assignment_1/internal/app"
)

func main() {
	var migrateFlag = flag.Bool("migrate", false, "Run the migration")
	flag.Parse()
	if *migrateFlag {
		app.RunMigrate()
	} else {
		app.Run()
	}
}
