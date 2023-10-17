package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func GetDB() (*gorm.DB, error) {
	return Db, nil
}
func Init(cfgFile string) {
	var err error

	ParseConfig(cfgFile)
	cfg := GetConfig().DbInfo
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("连接数据库失败,", err.Error())
		os.Exit(1)
	}

	if err != nil {
		fmt.Print("连接数据库失败2,", err.Error())
		os.Exit(2)
	}
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(cfg.DbMaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.DbMaxOpenConns)
}
