package app

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/nordew/tatl-test/internal/config"
	v1 "github.com/nordew/tatl-test/internal/http/controller/v1"
	"github.com/nordew/tatl-test/internal/service"
	"github.com/nordew/tatl-test/internal/storage"
	"github.com/nordew/tatl-test/pkg/db/mysql"
	"github.com/pressly/goose"
	"gorm.io/gorm"
)

func MustRun() {
	cfg := config.MustLoad()

	db, err := mysql.NewMySQLDB(cfg.MySQL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	migrate(db)

	logger := slog.Default()

	userStorage := storage.NewUserStorage(db)
	authStorage := storage.NewAuthStorage(db)

	userService := service.NewUserService(userStorage, logger)
	authService := service.NewAuthService(authStorage, logger)

	handler := v1.NewHandler(userService, authService)

	app := fiber.New()

	handler.RegisterRoutes(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%d", cfg.HttpRouter.Port)
		logger.Info("Server is starting", "address", addr)
		if err := app.Listen(addr); err != nil {
			logger.Error("Error starting server", "error", err)
		}
	}()

	<-quit
	logger.Info("shutting down server...")

	if err := app.Shutdown(); err != nil {
		logger.Error("error during server shutdown", "error", err)
	}
}

func migrate(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	goose.SetDialect("mysql")
	migrationsDir := "migrations"

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
