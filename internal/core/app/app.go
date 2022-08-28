package app

import (
	"context"
	"fmt"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"github.com/reinanhs/golang-web-api-structure/internal/core/http"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

type AppInterface interface {
	GetVersion() string
	GetName() string
	GetContext() context.Context
	Run()
}

type App struct {
	name    string
	version string
	ctx     context.Context
}

var lock = &sync.Mutex{}

var (
	appInstance *App
)

func getInstance(ctx context.Context) *App {
	if appInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if appInstance == nil {
			fmt.Println("Creating app single instance now.")
			appInstance = New(ctx)
		} else {
			fmt.Println("Single app instance already created.")
		}
	} else {
		fmt.Println("Single app instance already created.")
	}

	return appInstance
}

func GetApp(ctx context.Context) *App {
	return getInstance(ctx)
}

func New(ctx context.Context) *App {

	app := &App{
		name:    ctx.Value("config").(*config.AppConfig).AppName,
		version: readVersionFromFile("VERSION"),
		ctx:     ctx,
	}

	return app
}

func readVersionFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines[0]
}

func (app App) GetVersion() string {
	return app.version
}

func (app App) GetName() string {
	return app.name
}

func (app App) GetContext() context.Context {
	return app.ctx
}

func (app App) Run() {
	app.GetContext().Value("server").(*http.Server).Run()
}
