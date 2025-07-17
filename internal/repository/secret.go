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

func (sr *SecretRepository) CreateSecret(name, username, note string, ciphertext, derivedKey []byte) (model.Secret, error) {
	secret := model.Secret{
		Name:       name,
		Username:   username,
		Note:       note,
		Ciphertext: ciphertext,
		DerivedKey: derivedKey,
	}
	sr.db.Create(&secret)
	return secret, nil
}

func (sr *SecretRepository) DeleteSecretById(id string) (model.Secret, error) {
	var secret model.Secret
	sr.db.Where("id = ?", id).First(&secret)
	sr.db.Delete(&secret)
	return secret, nil
}
