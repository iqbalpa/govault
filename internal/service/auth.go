package service

import (
	"fmt"
	"govault/internal/model"
	"govault/internal/repository"
)

type AuthService struct {
	r repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{
		r: r,
	}
}

func (as *AuthService) InitMasterPass(masterPass string) (model.Auth, error) {
	// hashed password
	if as.r.IsInitialized() {
		return model.Auth{}, fmt.Errorf("you've initialized the vault")
	}
	a, _ := as.r.InitMasterPass(masterPass)
	return a, nil
}
