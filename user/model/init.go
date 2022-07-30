package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

var DB *gorm.DB

func Database() {
	db, err := gorm.Open("mysql", viper.GetString("mysql.endpoint"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	mysqlDb := db.DB()
	mysqlDb.SetMaxIdleConns(20)
	mysqlDb.SetMaxOpenConns(100)
	mysqlDb.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migrate()
}
