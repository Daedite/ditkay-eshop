package gorm

import (
	"fmt"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"
	logger2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBConnector struct {
	Config config.DBConfig
}

func NewPostgresDBConnector(config config.DBConfig) *PostgresDBConnector {
	return &PostgresDBConnector{
		Config: config,
	}
}

var _ DBConnector = &PostgresDBConnector{}

func (p *PostgresDBConnector) Connect() (*gorm.DB, error) {
	logger, err := logger2.CreateGormLogger(p.Config.LogLevel())
	if err != nil {
		return nil, err
	}
	return gorm.Open(postgres.Open(p.getConnectionString()), &gorm.Config{
		Logger: logger,
	})
}

func (p *PostgresDBConnector) getConnectionString() string {
	config := p.Config
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host(),
		config.Name(),
		config.Password(),
		config.Name(),
		config.Port(),
		config.SSLMode(),
		config.TimeZone())

}
