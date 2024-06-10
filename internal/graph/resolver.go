package graph

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/otogenavi-backend/internal/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ResolverDeps
}

type ResolverDeps struct {
	fx.In

	Ent *ent.Client
}
