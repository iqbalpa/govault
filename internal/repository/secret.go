package repository

import (
	"govault/internal/model"

	"gorm.io/gorm"
)

const (
	IdKey       string = "id"
	NameKey     string = "name"
	UsernameKey string = "username"
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
	res := sr.db.Select(IdKey, NameKey, UsernameKey).Find(&secrets)
	if res.Error != nil {
		return []model.Secret{}, res.Error
	}
	return secrets, nil
}

func (sr *SecretRepository) GetSecretById(id string) (model.Secret, error) {
	var secret model.Secret
	res := sr.db.Where("id = ?", id).First(&secret)
	if res.Error != nil {
		return model.Secret{}, res.Error
	}
	return secret, nil
}

func (sr *SecretRepository) CreateSecret(name, username, note string, ciphertext, salt []byte) (model.Secret, error) {
	secret := model.Secret{
		Name:       name,
		Username:   username,
		Note:       note,
		Ciphertext: ciphertext,
		Salt:       salt,
	}
	res := sr.db.Create(&secret)
	if res.Error != nil {
		return model.Secret{}, res.Error
	}
	return secret, nil
}

func (sr *SecretRepository) DeleteSecretById(id string) (model.Secret, error) {
	var secret model.Secret
	sr.db.Where("id = ?", id).First(&secret)
	res := sr.db.Delete(&secret)
	if res.Error != nil {
		return model.Secret{}, res.Error
	}
	return secret, nil
}
