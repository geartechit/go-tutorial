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
	"net/http"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig()
	zaplogger := logger.NewZapLogger(cfg)
	dtovalidator, err := validator.NewDTOValidator()
	if err != nil {
		zaplogger.Error("error creating dto validator")
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
	employeeHdr := handlers.NewEmployeeHandler(employeeSvc, zaplogger, dtovalidator)

	departmentRepo := repositories.NewDepartmentRepository(q, db.Pool)
	departmentSvc := services.NewDepartmentService(departmentRepo, zaplogger)
	departmentHdr := handlers.NewDepartmentHandler(departmentSvc, zaplogger, dtovalidator)

	r := router.New(employeeHdr, departmentHdr)

	zaplogger.Info("starting server", logger.Field{Key: "server port", Value: cfg.Server.Port})
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		zaplogger.Error("failed to start server")
	}
}
