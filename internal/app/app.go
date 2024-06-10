package app

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/otogenavi-backend/internal/app/appconfig"
	"github.com/GalvinGao/otogenavi-backend/internal/app/appcontext"
	"github.com/GalvinGao/otogenavi-backend/internal/controller"
	"github.com/GalvinGao/otogenavi-backend/internal/infra"
	"github.com/GalvinGao/otogenavi-backend/internal/server"
	"github.com/GalvinGao/otogenavi-backend/internal/service"
	"github.com/GalvinGao/otogenavi-backend/internal/x/logger"
	"github.com/GalvinGao/otogenavi-backend/internal/x/logger/fxlogger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		controller.Module(),
		infra.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
