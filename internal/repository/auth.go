package repository

import (
	"govault/internal/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) InitMasterPass(hashedPass string) (model.Auth, error) {
	auth := model.Auth{
		HashedMasterPassword: hashedPass,
	}
	ar.db.Create(&auth)
	return auth, nil
}

func (ar *AuthRepository) IsInitialized() bool {
	var auth model.Auth
	res := ar.db.First(&auth)
	return res.Error == nil
}
