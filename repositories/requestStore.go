package repositories

import (
	"Technopark_3_Security/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type RequestStore struct {
	DB *gorm.DB
}

func (store *RequestStore) Set(req models.Request) error {
	if err := store.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (store *RequestStore) Get(limit int) ([]models.Request, error) {
	var requests []models.Request
	if err := store.DB.Limit(limit).Order("id desc").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (store *RequestStore) GetByID(id int) (models.Request, error) {
	var request models.Request
	if err := store.DB.Where("id = ?", id).First(&request).Error; err != nil {
		return models.Request{}, err
	}
	return request, nil
}

func CreatePostgresDB() (*RequestStore, error) {
	postgresClient, err := gorm.Open(postgres.Open("host=db user=anton password=db_password dbname=db_repeater port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = postgresClient.AutoMigrate(&models.Request{}); err != nil {
		return nil, err
	}

	return &RequestStore{DB: postgresClient}, nil
}
