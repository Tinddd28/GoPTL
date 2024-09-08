package main

import (
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/config"
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	log_ := logger.SetupLogger(cfg.Env)
	log_.Info("starting server", slog.String("env", cfg.Env))
	log_.Debug("debugging")
	log_.Info("cfg now: ", slog.Any("cfg", cfg))

	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	log_.Info("test", slog.String("dbURL", dbURL))
	log_.Info("type of cfg", slog.Any("data", fmt.Sprintf("%T, %T, %T", cfg.Database.User, cfg.Database.Password, cfg.Database.Host)))

	m, err := migrate.New(
		"file://internal/migrations",
		dbURL,
	)
	if err != nil {
		log_.Error("failed to create migration", slog.String("error", err.Error()))
	}

	if err := m.Up(); err != nil {
		log_.Error("failed to run migration", slog.String("error", err.Error()))
		return
	}

	log_.Info("migration completed")
}
