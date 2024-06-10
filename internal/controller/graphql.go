package controller

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/fx"

	"github.com/GalvinGao/otogenavi-backend/internal/ent"
	"github.com/GalvinGao/otogenavi-backend/internal/graph"
	"github.com/GalvinGao/otogenavi-backend/internal/middleware"
	"github.com/GalvinGao/otogenavi-backend/internal/x/fiberctx"
)

type GraphQL struct {
	fx.In

	Ent          *ent.Client
	ResolverDeps graph.ResolverDeps
	Middleware   middleware.Middleware
	Route        fiber.Router `name:"root"`
}

var tracer = otel.Tracer("graphql")

func operationName(ctx context.Context) string {
	operationContext := graphql.GetOperationContext(ctx)
	requestName := "nameless-operation"
	if operationContext.Doc != nil && len(operationContext.Doc.Operations) != 0 {
		op := operationContext.Doc.Operations[0]
		if op.Name != "" {
			requestName = op.Name
		}
	}

	return requestName
}

func RegisterGraphQL(c GraphQL) {
	graphConfig := graph.Config{
		Resolvers: &graph.Resolver{
			ResolverDeps: c.ResolverDeps,
		},
	}

	srv := handler.New(graph.NewExecutableSchema(graphConfig))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.FixedComplexityLimit(200))
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.Use(entgql.Transactioner{TxOpener: c.Ent})
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		tracerCtx, span := tracer.Start(fiberctx.UserCtx(ctx), operationName(ctx))
		defer span.End()

		fiberctx.FiberCtx(ctx).SetUserContext(tracerCtx)

		oc := graphql.GetOperationContext(ctx)
		span.SetAttributes(
			attribute.String("request.query", oc.RawQuery),
		)
		complexityExtension := "ComplexityLimit"
		complexityStats, ok := oc.Stats.GetExtension(complexityExtension).(*extension.ComplexityStats)
		if !ok {
			// complexity extension is not used
			complexityStats = &extension.ComplexityStats{}
		}

		if complexityStats.ComplexityLimit > 0 {
			span.SetAttributes(
				attribute.Int("request.complexityLimit", complexityStats.ComplexityLimit),
				attribute.Int("request.operationComplexity", complexityStats.Complexity),
			)
		}

		for key, val := range oc.Variables {
			span.SetAttributes(
				attribute.String(fmt.Sprintf("request.variables.%s", key), fmt.Sprintf("%+v", val)),
			)
		}

		return next(ctx)
	})

	c.Route.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/graphql")))
	c.Route.Post("/graphql", c.Middleware.InjectFiberCtx(), adaptor.HTTPHandler(srv))
}
