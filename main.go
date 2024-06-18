package main

import (
	"context"

	"github.com/webdevelop-pro/go-common/configurator"
	"github.com/webdevelop-pro/go-common/db"
	"github.com/webdevelop-pro/go-common/logger"
	"github.com/webdevelop-pro/go-common/verser"
	"github.com/webdevelop-pro/i-models/profiles"
	"go.uber.org/fx"
)

var (
	//nolint:gochecknoglobals
	service string
	version string
	//nolint:gochecknoglobals
	repository string
	//nolint:gochecknoglobals
	revisionID string
)

func main() {
	verser.SetServiVersRepoRevis(service, version, repository, revisionID)
	log := logger.NewComponentLogger(context.Background(), "fx")
	fx.New(
		fx.Logger(log),
		fx.Provide(
			// Application
			context.Background,
			configurator.NewConfigurator,
			// Storages
			db.New,
		),

		fx.Invoke(
			// Run HTTP server
			Hello,
		),
	).Run()
}

func Hello(lc fx.Lifecycle, pool *db.DB) {
	ctx := context.Background()
	log := logger.NewComponentLogger(ctx, "fx")
	profile := profiles.New(pool)
	profile, err := profile.SelectByID(ctx, 123)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get profile")
	}
}
