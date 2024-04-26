package config

import (
	"fmt"
	"net/url"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMysql(env string, cfg *MysqlOption) (*Mysql, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		url.QueryEscape(cfg.TimeZone),
	)

	logLevel := glogger.Warn
	if env == "local" {
		logLevel = glogger.Info
	}

	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(cfg.Pool)
	return &Mysql{DB: db}, err
}
