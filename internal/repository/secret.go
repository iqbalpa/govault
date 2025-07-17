package repository

import (
	"govault/internal/model"

	"gorm.io/gorm"
)

const (
	IdKey   string = "id"
	NameKey string = "name"
)

type SecretRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SecretRepository {
	return &SecretRepository{
		db: db,
	}
}

func (sr *SecretRepository) GetAllSecrets() ([]model.Secret, error) {
	var secrets []model.Secret
	sr.db.Select(IdKey, NameKey).Find(&secrets)
	return secrets, nil
}

func (sr *SecretRepository) GetSecretById(id string) (model.Secret, error) {
	var secret model.Secret
	sr.db.Where("id = ?", id).First(&secret)
	return secret, nil
}
