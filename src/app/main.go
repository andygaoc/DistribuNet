package main

import (
	"app/config"
	"app/models"
	_ "gorm.io/gorm"
	"log"
)

func main() {
	db, err := config.GetDB()
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建数据库表结构
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		panic("failed to migrate database")
	}

	// 在此处编写其他代码...
	// 关闭数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
