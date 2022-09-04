package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/core/config"
	"github.com/reinanhs/golang-web-api-structure/internal/http/middleware"
	"github.com/reinanhs/golang-web-api-structure/internal/router"
	"log"
)

type ServerInterface interface {
	GetHost() string
	GetPort() int
	GetContext() context.Context
	Run()
}

// Server represents a http server that listens on a port.
type Server struct {
	host   string
	port   string
	server *gin.Engine
	ctx    context.Context
}

// New instantiates a new instance of Server.
func New(ctx context.Context) *Server {
	appConfig := ctx.Value("config").(*config.AppConfig)

	return &Server{
		host:   appConfig.AppHost,
		port:   appConfig.AppPort,
		server: gin.Default(),
		ctx:    ctx,
	}
}

// Run the server
func (s *Server) Run() {
	s.server.Use(middleware.HandleCORS(s.server))
	s.server.Use(middleware.TranslationMiddleware())

	router.InitRouter(s.ctx, s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Printf("Link: %s:%v", s.host, s.port)
	log.Fatal(s.server.Run(":" + s.port))
}

func (s *Server) GetHost() string {
	return s.host
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) GetContext() context.Context {
	return s.ctx
}
