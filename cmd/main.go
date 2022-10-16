package main

import (
	"github.com/abdukhashimov/golang-hex-architecture/api"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/logger"
	"github.com/abdukhashimov/golang-hex-architecture/service"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
)

func main() {
	logger := logger.NewLogger()
	cfg := config.Load()

	strg := storage.NewStorage(nil)
	services := service.NewServiceHandler(&cfg, logger, strg)

	server := api.New(&api.RouterOptions{
		Cfg:     &cfg,
		Log:     logger,
		Service: services,
	})

	server.Run(cfg.HTTPPort)
}
