package main

import (
	"context"
	"fmt"
	"github.com/felipedenardo/chameleon-auth-api/internal/app"
	"github.com/felipedenardo/chameleon-auth-api/internal/config"
	"github.com/felipedenardo/chameleon-auth-api/internal/infra/database/postgresql/migration"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// @title Auth API Microservice (Chameleon System)
// @version 1.0
// @description Este serviço é o Provedor Central de Identidade (IAM)...
// @host localhost:8081
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Digite o token no formato Bearer {token}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] No .env file found, using system environment variables.")
	}
	cfg := config.Load()

	log.Printf("Starting Auth API on port %s...", cfg.Port)

	db := setupPostgres(cfg)
	redisClient := setupRedis(cfg)

	handlers := app.NewHandlerContainer(db, cfg, redisClient)
	r := app.SetupRouter(handlers, cfg)

	log.Printf("[INFO] Auth API running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("[FATAL] Server failed: %v", err)
	}
}

func setupPostgres(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[FATAL] Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("PostgreSQL connection established")

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migration.ID011220251300DDLCreateInitialSchema,
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("[FATAL] Migration failed: %v", err)
	}
	log.Println("Migrations executed successfully")

	return db
}

func setupRedis(cfg *config.Config) *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("[FATAL] Failed to connect to Redis: %v", err)
	}
	log.Println("Redis connection established")

	return redisClient
}
