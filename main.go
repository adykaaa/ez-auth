package main

import (
	"adykaaa/ez-auth/config"
	"adykaaa/ez-auth/db"
	"adykaaa/ez-auth/logger"
	"log"

	"github.com/rs/zerolog"
)

func main() {
	l := logger.New(zerolog.InfoLevel.String())

	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	_, err = db.NewSQL("postgres", config.DBConnString, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = db.MigrateDB(config.DBConnString, "file://db/migrations/", &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

}
