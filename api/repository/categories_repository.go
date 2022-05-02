package repository

import (
	"project/api/models"

	"gorm.io/gorm"
)

type CategoriesRepository interface{
	Save(*models.Category) (*models.Category,error)
}

type categoriesRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *categoriesRepositoryImpl {
	return &categoriesRepositoryImpl{db}
}

func (r *categoriesRepositoryImpl) Save(category *models.Category) (*models.Category, error) {


	tx := r.db.Debug().Begin()
	err := tx.Debug().Model(&models.Product{}).Create(category).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return category, tx.Commit().Error

}