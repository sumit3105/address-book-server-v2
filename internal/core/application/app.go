package application

import (
	logger "address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/core/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Cfg *config.Config
	DB  *gorm.DB
}

func NewApp() *App {
	// Init logger
	logger.InitLogger()
	defer logger.Logger.Sync()

	// Load config
	cfg := load()

	// Connect DB and Migrate
	db := connect(cfg)
	// a.DB.AutoMigrate(&models.User{}, &models.Address{})

	return &App{
		Cfg: cfg,
		DB:  db,
	}
}

func load() *config.Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Warn("No .env file found, using system environment variables")

	}

	dbCfg := &config.DBConfig{
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
	}

	serverCfg := &config.ServerConfig{
		ServerPort: getEnv("SERVER_PORT"),
		JwtSecret:  getEnv("JWT_SECRET"),
		AppURL:     getEnv("APP_URL"),
	}

	smtpCfg := &config.SMTPConfig{
		SMTPUser: getEnv("SMTP_USER"),
		SMTPPass: getEnv("SMTP_PASS"),
		SMTPHost: getEnv("SMTP_HOST"),
		SMTPPort: getEnv("SMTP_PORT"),
	}

	cfg := &config.Config{
		DBCfg:     dbCfg,
		ServerCfg: serverCfg,
		SMTPCfg:   smtpCfg,
	}

	logger.Logger.Info("config loaded")
	return cfg
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		logger.Logger.Fatal("environment variable required", zap.String("key", key))

	}
	return value
}

func connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBCfg.DBHost,
		cfg.DBCfg.DBPort,
		cfg.DBCfg.DBUser,
		cfg.DBCfg.DBPassword,
		cfg.DBCfg.DBName,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("failed to connect to database", zap.Error(err))
	}

	logger.Logger.Info("database connected")

	return database
}
