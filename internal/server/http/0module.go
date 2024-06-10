package http

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/otogenavi-backend/internal/server/http/route"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), route.Module())
}
