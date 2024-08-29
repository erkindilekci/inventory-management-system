package app

import (
	"ims-intro/pkg/common/postgresql"
	"os"
)

type ConfigurationManager struct {
	PostgresqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	postgresqlConfig := postgresql.Config{
		Host:                  dbHost,
		Port:                  dbPort,
		UserName:              dbUser,
		Password:              dbPassword,
		DbName:                dbName,
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
	return &ConfigurationManager{postgresqlConfig}
}
