package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/BoomTHDev/wear-pos-server/config"
	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type fiberServer struct {
	app  *fiber.App
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *fiberServer
)

func NewFiberServer(conf *config.Config, db databases.Database) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:   conf.Server.BodyLimit,
		IdleTimeout: time.Second * time.Duration(conf.Server.TimeOut),
	})

	once.Do(func() {
		server = &fiberServer{
			app:  fiberApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *fiberServer) Start() {
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%v", s.conf.Server.AllowOrigins[0]),
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	s.app.Get("/v1/health", s.healthCheck)

	s.httpListening()
}

func (s *fiberServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Listen(url); err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func (s *fiberServer) healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("OK")
}
