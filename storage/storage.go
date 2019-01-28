package storage

import (
	"Mock-API-Data/config"
	"Mock-API-Data/model"
	"errors"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Storage struct {
	conf *config.Config
	db   *gorm.DB
}

func NewStorage(conf *config.Config) (*Storage, error) {
	if conf == nil {
		return nil, errors.New("conf can be not nil")
	}
	if conf.DBPath == "" {
		return nil, errors.New("DBPath can be not empty")
	}
	db, err := gorm.Open("sqlite3", conf.DBPath)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// 自动建表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.Data{})
	db.AutoMigrate(&model.Rule{})
	s := &Storage{
		conf: conf,
		db:   db,
	}
	return s, nil
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
