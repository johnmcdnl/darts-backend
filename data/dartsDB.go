package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

var once sync.Once
var db *gorm.DB

func GetDB() *gorm.DB {
	once.Do(func() {
		initDb()
	})
	return db
}

func initDb() {
	conn, err := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = &conn
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
}
