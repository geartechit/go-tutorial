package main

import (
	"context"
	"go-tutorial/internal/config"
	"go-tutorial/internal/database"
	"go-tutorial/internal/database/sqlc/queries"
	"go-tutorial/internal/handlers"
	"go-tutorial/internal/repositories"
	"go-tutorial/internal/router"
	"go-tutorial/internal/services"
	"go-tutorial/pkg/logger"
	"go-tutorial/pkg/validator"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig()
	applogger := logger.New(cfg)
	zaplogger := logger.NewZapLogger(cfg)
	dtovalidator, err := validator.NewDTOValidator()
	if err != nil {
		applogger.Error("error creating dto validator", zap.Error(err))
		return
	}

	db, err := database.Connect(ctx, cfg.Postgres)
	if err != nil {
		panic(err)
	}
	defer db.Pool.Close()
	q := queries.New(db.Pool)

	employeeRepo := repositories.NewEmployeeRepository(q)
	employeeSvc := services.NewEmployeeService(employeeRepo, zaplogger)
	employeeHdr := handlers.NewEmployeeHandler(employeeSvc, applogger, dtovalidator)

	r := router.New(employeeHdr)

	applogger.Info("starting server", zap.String("port", cfg.Server.Port))
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		applogger.Error("failed to start server", zap.Error(err))
	}
}
