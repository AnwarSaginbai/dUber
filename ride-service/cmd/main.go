package main

import (
	"database/sql"
	"fmt"
	"github.com/AnwarSaginbai/ride-service/internal/adapter/grpc"
	"github.com/AnwarSaginbai/ride-service/internal/adapter/postgres"
	"github.com/AnwarSaginbai/ride-service/internal/config"
	"github.com/AnwarSaginbai/ride-service/internal/logging"
	"github.com/AnwarSaginbai/ride-service/internal/service"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	logger := logging.InitLogging(cfg.Env)
	db, err := OpenDatabase(
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}
	logger.Info("successfully connected to database")
	store := postgres.NewPostgres(db)
	api := service.NewService(store)
	app := grpc.NewAdapter(api, logger, cfg)
	logger.Info("starting the ride application on addr",
		"addr",
		fmt.Sprintf(
			"%s:%d",
			cfg.Server.Host,
			cfg.Server.Port), "env-level", cfg.Env,
	)

	if err = app.Run(); err != nil {
		logger.Error("Failed to start ride-service", "error", err)
	}
}

func OpenDatabase(username, password, host, port, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
