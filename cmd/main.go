package main

import (
	"github.com/Tinddd28/GoPTL/internal/config"
	"github.com/Tinddd28/GoPTL/internal/handler"
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/internal/service"
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/Tinddd28/GoPTL/pkg/server"
	"log/slog"
	"time"
)

// @title GoPTL API
// @version 1
// @description This is a backend for PrimeTokenList.

// @host 0.0.0.0:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	time.Sleep(time.Second)
	cfg := config.MustLoad()
	log_ := logger.SetupLogger(cfg.Env)
	log_.Info("starting server", slog.String("env", cfg.Env))
	log_.Debug("debugging")
	log_.Info("cfg now: ", slog.Any("cfg", cfg))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Username: cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  "disable",
	})
	if err != nil {
		log_.Error("failed to initialize db", slog.String("error", err.Error()))
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	if err := srv.Run(cfg.HTTPServer.Address, handlers.InitRoutes()); err != nil {
		log_.Error("server run failed", slog.String("error", err.Error()))
	}

}
