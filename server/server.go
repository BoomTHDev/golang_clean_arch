package server

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/BoomTHDev/golang_clean_arch/config"
	"github.com/BoomTHDev/golang_clean_arch/databases"
	"github.com/BoomTHDev/golang_clean_arch/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app   *fiber.App
	db    databases.Database
	conf  *config.Config
	redis databases.RedisClient
}

var (
	once           sync.Once
	serverInstance *fiberServer
)

func NewFiberServer(conf *config.Config, db databases.Database, redis databases.RedisClient) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:    conf.Server.BodyLimit,
		IdleTimeout:  time.Second * time.Duration(conf.Server.TimeOut),
		ErrorHandler: middleware.ErrorHandler(),
	})

	once.Do(func() {
		serverInstance = &fiberServer{
			app:   fiberApp,
			db:    db,
			conf:  conf,
			redis: redis,
		}
	})

	return serverInstance
}

func (s *fiberServer) setupRoutes() {
	s.app.Use(logger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(s.conf.Server.AllowOrigins, ","),
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	s.app.Get("/v1/health", s.healthCheck)
	s.initUserRouter()

	s.app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Sorry, endpoint %s %s not found", ctx.Method(), ctx.Path()),
		})
	})
}

func (s *fiberServer) Start() {
	s.setupRoutes()
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
