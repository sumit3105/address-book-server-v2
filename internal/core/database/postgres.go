package database

import (
	"address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/core/config"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
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