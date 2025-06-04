package databasex

import "time"

const (
	connStringMysqlTemplate = "%s:%s@tcp(%s:%d)/%s?%s"
	//connStringPostgresTemplate = "%s:%s@%s:%d/%s?%s"
	connStringPostgresTemplate = "postgres://%s/%s?%s"
)

type (
	Config struct {
		Host         string
		Port         int
		User         string
		Password     string
		Name         string
		Charset      string
		MaxOpenConns int
		MaxIdleConns int
		MaxLifetime  time.Duration
		Type         string
		TimeZone     string
		IsEncrypt    bool
		Driver       string
		SSLServer    string
		DialTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)
