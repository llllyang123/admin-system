package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB 定义数据库连接接口
type DB interface {
	Exec(sql string, values ...interface{}) (*gorm.DB, error)
	Query(sql string, values ...interface{}) (*gorm.DB, error)
	// ... 其他需要的方法
}

type gormDB struct {
	*gorm.DB
}

func NewGormDB(dsn string) (*gormDB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &gormDB{db}, nil
}
