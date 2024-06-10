package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/GalvinGao/otogenavi-backend/internal/x/fiberctx"
)

type Middleware struct {
	fx.In
}

func (m Middleware) InjectFiberCtx() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// inject fiber context into *fasthttp.RequestCtx
		ctx.Context().SetUserValue(fiberctx.CtxKeyFiberCtx, ctx)
		return ctx.Next()
	}
}
