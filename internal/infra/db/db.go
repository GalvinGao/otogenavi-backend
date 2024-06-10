package db

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/GalvinGao/otogenavi-backend/internal/app/appconfig"
	"github.com/GalvinGao/otogenavi-backend/internal/ent"
)

func New(conf *appconfig.Config) *ent.Client {
	client, err := ent.Open("postgres", conf.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed connecting to database")
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}
	return client.Debug()
}
