package services

import (
	"app/config"
	"app/models"
)

func CreateProduct(product *models.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetProductByID(id string) (*models.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var product models.Product
	result := db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// 其他服务函数...
