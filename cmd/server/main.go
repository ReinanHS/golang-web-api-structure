package main

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/core/app"
)

//Execution starts from main function
func main() {
	ctx := context.Background()
	ctx = app.GetContainer(ctx)

	//ctx = context.WithValue(ctx, "config", config.GetConfig(ctx))
	//ctx = context.WithValue(ctx, "server", http.New(ctx))
	//ctx = context.WithValue(ctx, "db", mysql.New(ctx))

	appInstance := app.GetApp(ctx)
	appInstance.Run()
}
