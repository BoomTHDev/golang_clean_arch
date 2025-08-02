package main

import (
	"github.com/BoomTHDev/golang_clean_arch/config"
	"github.com/BoomTHDev/golang_clean_arch/databases"
	"github.com/BoomTHDev/golang_clean_arch/server"
)

func main() {
	cfg := config.ConfigGetting()
	db := databases.NewPostgresDatabase(cfg.Database)
	redis := databases.NewRedisClient(cfg.Redis)
	server := server.NewFiberServer(cfg, db, redis)

	server.Start()
}
