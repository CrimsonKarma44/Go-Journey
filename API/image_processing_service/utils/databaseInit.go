package utils

import (
	auth_model "IPS/auth/model"
	img_model "IPS/img/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(values auth_model.Data_values) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", values.Host, values.User, values.Password, values.Db_Name, values.Ports)), &gorm.Config{})
	if err != nil {

		return nil, fmt.Errorf("error occured")
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&auth_model.Creds{}, &img_model.Image{})
	if err != nil {
		return err
	}
	return nil
}
