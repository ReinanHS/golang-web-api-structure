package main

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/core/app"
)

//Execution starts from main function
func main() {
	ctx := context.Background()
	ctx = app.GetContainer(ctx)

	appInstance := app.GetApp(ctx)
	appInstance.RunMigration()
	appInstance.Run()
}
