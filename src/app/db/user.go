package db

import (
	"app/models"
)

func CreateUser(user *models.User) error {
	db := GetDB()

	// 执行插入操作

	return nil
}

func GetUserByID(id int) (*models.User, error) {
	db := GetDB()

	// 执行查询操作

	return nil, nil
}

// 其他数据库操作函数...
