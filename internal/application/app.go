package application

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teesh3rt/wizardapi/internal/database"
)

type App struct {
	Router  *fiber.App
	Queries *database.Queries
	DB      *pgxpool.Pool
}

type dbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func New() App {
	router := fiber.New()
	router.Use(logger.New(), recover.New())

	cfg := dbConfig{
		Host:     getEnv("DATABASE_HOST", "localhost"),
		Port:     getEnv("DATABASE_PORT", "5432"),
		Name:     getEnv("DATABASE_DB", "postgres"),
		User:     getEnv("DATABASE_USER", "postgres"),
		Password: getEnv("DATABASE_PASSWORD", "postgres"),
	}

	dbUrl := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Name + "?sslmode=disable"

	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	queries := database.New(dbpool)

	app := App{
		Router:  router,
		Queries: queries,
		DB:      dbpool,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start() {
	a.Router.Listen(":8080")
}
