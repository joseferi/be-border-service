package migration

import (
	"be-border-service/internal/config"
	"be-border-service/pkg/databasex"
)

func MigrationDatabase() {
	config := config.RegisterConfiguration()
	databasex.DatabaseMigration(&databasex.Config{
		Host:         config.DB.Host,
		Port:         config.DB.Port,
		User:         config.DB.User,
		Password:     config.DB.Password,
		Name:         config.DB.Name,
		IsEncrypt:    config.DB.IsEncrypt,
		MaxOpenConns: config.DB.MaxOpenConns,
		MaxIdleConns: config.DB.MaxIdleConns,
		MaxLifetime:  config.DB.MaxLifetime,
		TimeZone:     config.DB.TimeZone,
		Driver:       config.DB.Driver,
		SSLServer:    config.DB.SSLServer,
		ReadTimeout:  config.DB.ReadTimeout,
		DialTimeout:  config.DB.DialTimeout,
		WriteTimeout: config.DB.WriteTimeout,
	})
}
