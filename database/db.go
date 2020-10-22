package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Initialize() {
	db, err := gorm.Open("mysql", viper.GetString("secrets.cloud-sql-user")+":"+viper.GetString("secrets.cloud-sql-password")+"@"+viper.GetString("cloud-sql.gormhost")+"/"+viper.GetString("cloud-sql.dbname")+"?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(viper.GetInt("cloud-sql.max-idle-conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("cloud-sql.max-open-conns"))
	db.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("cloud-sql.max-life-time-minutes")) * time.Minute)
	if err != nil {
		log.Fatal("Init DB Error", err)
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func Close() {
	DB.Close()
}
