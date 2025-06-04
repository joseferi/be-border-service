package databasex

import (
	"context"
	"fmt"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	connStringMssqlTemplate = "sqlserver://%s:%s@%s:%v?database=%s&log=3"
)

func mssqlDSN(cfg *Config) string {
	dsn := fmt.Sprintf(connStringMssqlTemplate,
		url.QueryEscape(cfg.User),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	if cfg.IsEncrypt {
		dsn = fmt.Sprintf(
			"%s&encrypt=%s&TrustServerCertificate=true&hostNameInCertificate=%s",
			url.QueryEscape(dsn),
			"true",
			cfg.SSLServer,
		)
	}
	return dsn
}
func NewMssql(cfg *Config) (*sqlx.DB, error) {

	db, err := sqlx.Open(cfg.Driver, mssqlDSN(cfg))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return db, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)
	return db, nil
}
