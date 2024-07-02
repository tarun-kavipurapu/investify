package main

import (
	"investify/config"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	config.LoadConfig(".")
	dbSource := config.EnvVars.DBSource
	log.Println(dbSource)

	m, err := migrate.New("file://db/migration", dbSource)

	if err != nil {
		log.Fatal("Init error", err.Error())
	}

	arg := os.Args[1]

	if arg == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration error", err)
		}

		log.Println("Migration down successful")
		return
	}

	if arg == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration error", err)
		}

		log.Println("Migration up successful")
		return
	}
}
