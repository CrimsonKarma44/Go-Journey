package services

import (
	"IPS/img/models"

	"gorm.io/gorm"
)

type ImageService struct {
	DB *gorm.DB
}

func (s ImageService) New(db *gorm.DB) ImageService {
	return ImageService{
		DB: db,
	}
}

func (s ImageService) SaveImage(filename string, path string) (error, uint) {
	image := models.Image{
		Filename: filename,
		Path:     path,
	}
	return s.DB.Create(&image).Error, image.ID
}

func (s ImageService) GetImageByID(id int) (*models.Image, error) {
	var image models.Image
	if err := s.DB.First(&image, id).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (s ImageService) GetImages(limit int, offset int) ([]models.Image, error) {
	var image []models.Image
	if err := s.DB.Offset(offset).Limit(limit).Find(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}
