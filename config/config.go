package config

import (
	"os"
	"strconv"
)

// Config define product-integrator config structure
type Config struct {
	AppEnv  string
	ApiPort int
	MysqlOption
}

// MysqlOption contains mySQL connection options
type MysqlOption struct {
	Driver       string
	Host         string
	Port         int
	Pool         int
	DatabaseName string
	Username     string
	Password     string
	TimeZone     string
}

// NewConfig initialize new config
func NewConfig() *Config {
	var cfg Config
	cfg.setPort().setEnvironment().setMySql()

	return &cfg
}

func (c *Config) setPort() *Config {
	v := os.Getenv("API_PORT")
	c.ApiPort, _ = strconv.Atoi(v)

	return c
}

func (c *Config) setEnvironment() *Config {
	c.AppEnv = os.Getenv(`APP_ENV`)
	return c
}

func (c *Config) setMySql() *Config {
	p := os.Getenv(`MYSQL_PORT`)
	mySqlPort, _ := strconv.Atoi(p)

	c.MysqlOption = MysqlOption{
		Driver:       os.Getenv(`MYSQL_DRIVER`),
		Host:         os.Getenv(`MYSQL_HOST`),
		Port:         mySqlPort,
		DatabaseName: os.Getenv(`MYSQL_DATABASE_NAME`),
		Username:     os.Getenv(`MYSQL_USERNAME`),
		Password:     os.Getenv(`MYSQL_PASSWORD`),
		TimeZone:     os.Getenv(`MYSQL_TIMEZONE`),
	}
	return c
}
