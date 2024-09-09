package main

import (
	"github.com/Tinddd28/GoPTL/internal/config"
	"github.com/Tinddd28/GoPTL/internal/handler"
	"github.com/Tinddd28/GoPTL/internal/repository"
	"github.com/Tinddd28/GoPTL/internal/service"
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/Tinddd28/GoPTL/pkg/server"
	"log/slog"
)

// @title GoPTL API
// @version 1
// @description This is a backend for PrimeTokenList.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
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

//func startServer(router *httprouter.Router, cfg *config.Config, log_ *slog.Logger) {
//	listener, err := net.Listen("tcp", fmt.Sprintf("%s", cfg.HTTPServer.Address))
//	if err != nil {
//		panic(err)
//	}
//	log_.Info("server listener", slog.String("port", cfg.HTTPServer.Address))
//	server := &http.Server{
//		Handler:      router,
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//
//	err = server.Serve(listener)
//	log.Fatalln(server.Serve(listener))
//
//}
