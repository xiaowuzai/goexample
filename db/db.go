package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	*gorm.DB
}

func NewDBLogger(env string) logger.Interface {
	level := logger.Warn
	if env == "offline" {
		level = logger.Info
	}
	return logger.Default.LogMode(level)
}

func NewDB(dsn string, l logger.Interface) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: l,
	})
	if err != nil {
		return nil, err
	}
	db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(10).
			SetMaxOpenConns(100),
	)

	return &DB{db}, nil
}
