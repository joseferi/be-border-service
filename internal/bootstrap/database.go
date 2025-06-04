package bootstrap

import (
	"be-border-service/internal/config"
	"be-border-service/pkg/databasex"
	"be-border-service/pkg/logger"

	"github.com/jmoiron/sqlx"
)

func RegistryDatabase(cfg *config.Database) *sqlx.DB {
	db, err := databasex.NewMssql(&databasex.Config{
		Host:         cfg.Host,
		Port:         cfg.Port,
		User:         cfg.User,
		Password:     cfg.Password,
		Name:         cfg.Name,
		Charset:      cfg.Charset,
		MaxOpenConns: cfg.MaxOpenConns,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxLifetime:  cfg.MaxLifetime,
		TimeZone:     cfg.TimeZone,
		IsEncrypt:    cfg.IsEncrypt,
		Driver:       cfg.Driver,
		SSLServer:    cfg.SSLServer,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
	if err != nil {
		logger.Fatal(err.Error(),
			logger.EventName("db"),
			logger.Any("host", cfg.Host),
			logger.Any("port", cfg.Port),
			logger.Any("driver", cfg.Driver),
			logger.Any("timezone", cfg.TimeZone),
		)
	}
	return db
}
