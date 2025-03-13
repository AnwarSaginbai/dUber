package main

import (
	"database/sql"
	"fmt"
	"github.com/AnwarSaginbai/auth-service/internal/adapter/grpc"
	"github.com/AnwarSaginbai/auth-service/internal/adapter/postgres"
	"github.com/AnwarSaginbai/auth-service/internal/config"
	"github.com/AnwarSaginbai/auth-service/internal/logging"
	"github.com/AnwarSaginbai/auth-service/internal/service"
	"log"
)

func main() {
	cfg := config.InitConfig()

	logger := logging.InitLogging(cfg.Env)

	db, err := OpenDatabase(
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	if err != nil {
		logger.Error("failed to open database", "src", "main.go")
		log.Fatalf("failed to open database: %v", err)
	}
	logger.Info("successfully connected to database")

	store := postgres.NewPostgresStorage(db)
	api := service.NewAPI(store)
	app := grpc.NewAdapter(cfg, api)

	logger.Info("starting the auth application on addr",
		"addr",
		fmt.Sprintf(
			"%s:%d",
			cfg.Server.Host,
			cfg.Server.Port), "env-level", cfg.Env,
	)

	if err = app.Run(); err != nil {
		logger.Error("failed to listen the server", "src", "main.go")
		log.Fatalf("error starting server: %v", err)
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
