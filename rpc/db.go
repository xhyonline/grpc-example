package rpc

import (
	"github.com/jinzhu/gorm"
	"github.com/xhyonline/grpc-example/mod"
	"github.com/xhyonline/xutil/db"
	"sync"
)

var once sync.Once

var instance *gorm.DB

func GetDB() *gorm.DB {
	once.Do(func() {
		instance = db.NewDataBase(&db.Config{
			Host:     "101.200.150.2",
			Port:     "3306",
			User:     "root",
			Password: "xxxxx",
			Name:     "grpc",
		})
		instance.AutoMigrate(&mod.User{}, &mod.Phone{})
	})
	return instance
}
