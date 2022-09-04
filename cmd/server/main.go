package main

import (
	"context"
	_ "github.com/reinanhs/golang-web-api-structure/docs"
	"github.com/reinanhs/golang-web-api-structure/internal/core/app"
)

// @title                     Golang web api structure
// @version                   1.0
// @description               A sample API project with golang
// @host                      localhost:8080
// @BasePath                  /api/v1
// @securityDefinitions.basic BasicAuth
func main() {
	ctx := context.Background()
	ctx = app.GetContainer(ctx)

	appInstance := app.GetApp(ctx)
	appInstance.RunMigration()
	appInstance.Run()
}
