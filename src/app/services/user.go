package services

import (
	"app/config"
	"app/models"
)

func CreateUser(user *models.User) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByID(id string) (*models.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// 其他服务函数...
