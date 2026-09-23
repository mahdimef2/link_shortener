package repository

import (
	"gorm.io/gorm"
	"url_shotener/internal/models"
)

type URLRepository interface {
	Create(url *models.Url) error
	FindByCode(code string) (*models.Url, error)
	UpdateCode(url *models.Url) error
}

type gormURLRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) URLRepository {
	return &gormURLRepository{db: db}
}

func (r *gormURLRepository) Create(url *models.Url) error {
	return r.db.Create(url).Error
}
func (r *gormURLRepository) FindByCode(code string) (*models.Url, error) {
	var url models.Url
	err := r.db.Where("code = ?", code).First(&url).Error
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *gormURLRepository) UpdateCode(url *models.Url) error {
	return r.db.Model(url).Update("code", url.Code).Error
}
