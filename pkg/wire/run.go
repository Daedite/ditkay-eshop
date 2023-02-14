package wire

import (
	eshop_api "github.com/ESPOIR-DITE/ditkay-eshop/pkg/config/eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/connect/gorm"
	migrate2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/migration/migrate"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
)

func Run(configManager eshop_api.ConfigManager) int {
	config, err := configManager.Load()
	if err != nil {
		logger.Log.Fatal("Failed to load application configurations: %s", err.Error())
		return 1
	}
	if err := logger.Configure(config.AppConfig().LogLevel()); err != nil {
		logger.Log.Fatal("Failed to configure logger: %s", err.Error())
	}
	gormDB, err := gorm.NewPostgresDBConnector(config.DBConfig()).Connect()
	if err != nil {
		logger.Log.Fatalf(err.Error())
		return 2
	}
	migrator, err := migrate2.NewPostgresMigrator(config.DBConfig()).NewMigrator()
	if err != nil {
		logger.Log.Fatalf("Faile to initiate db migration: %s", err.Error())
		return 3
	}
	err = migrator.Up()
	if err != migrate.ErrNoChange && err != nil {
		logger.Log.Fatalf("Faile to run migration: %s", err)
		return 4
	}

	httpServer := wire(config, gormDB)

	if err = httpServer.Start(); err != nil {
		logger.Log.Fatal(err)
		return 1
	}
	return 0
}
