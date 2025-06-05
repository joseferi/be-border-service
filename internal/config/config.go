package config

import "time"

var (
	Cfg Config
)

type Config struct {
	Server Server      `mapstructure:",squash"`
	DB     Database    `mapstructure:",squash"`
	Redis  Redis       `mapstructure:",squash"`
	Asynq  AsynqServer `mapstructure:",squash"`
}

type Server struct {
	Name string `mapstructure:"APP_NAME"`
	Port int    `mapstructure:"APP_PORT"`
	Env  string `mapstructure:"APP_ENV"`
}
type AsynqServer struct {
	Concurrency int `mapstructure:"ASYNQ_SERVER_CONCURRENCY"`
}

type Database struct {
	Host         string        `mapstructure:"DB_HOST"`
	Port         int           `mapstructure:"DB_PORT"`
	User         string        `mapstructure:"DB_USER"`
	Password     string        `mapstructure:"DB_PASSWORD"`
	Name         string        `mapstructure:"DB_NAME"`
	Charset      string        `mapstructure:"DB_CHARSET"`
	MaxOpenConns int           `mapstructure:"DB_MAX_OPEN_CONN"`
	MaxIdleConns int           `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxLifetime  time.Duration `mapstructure:"DB_MAX_LIFETIME"`
	TimeZone     string        `mapstructure:"DB_TIMEZONE"`
	IsEncrypt    bool          `mapstructure:"DB_IS_ENCRYPT"`
	Driver       string        `mapstructure:"DB_DRIVER"`
	SSLServer    string        `mapstructure:"DB_SSL_SERVER"`
	DialTimeout  time.Duration `mapstructure:"DB_DIAL_TIMEOUT"`
	ReadTimeout  time.Duration `mapstructure:"DB_READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"DB_WRITE_TIMEOUT"`
}

type Redis struct {
	Addr         string        `mapstructure:"REDIS_ADDR"`
	Username     string        `mapstructure:"REDIS_USERNAME"`
	Password     string        `mapstructure:"REDIS_PASSWORD"`
	DB           int           `mapstructure:"REDIS_DB"`
	ReadTimeout  time.Duration `mapstructure:"REDIS_READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"REDIS_WRITE_TIMEOUT"`
}
