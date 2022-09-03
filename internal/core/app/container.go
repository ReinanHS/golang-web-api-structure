package app

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"github.com/reinanhs/golang-web-api-structure/internal/core/drivers"
	"github.com/reinanhs/golang-web-api-structure/internal/core/http"
)

func GetContainer(ctx context.Context) context.Context {

	// App Core
	ctx = context.WithValue(ctx, "config", config.GetConfig())
	ctx = context.WithValue(ctx, "db", drivers.New(ctx))
	ctx = context.WithValue(ctx, "server", http.New(ctx))

	return ctx
}
