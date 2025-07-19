package service

import (
	"fmt"
	"govault/internal/model"
	"govault/internal/repository"
	"govault/internal/utils"
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
	hashed, _ := utils.HashPassword(masterPass)
	if as.r.IsInitialized() {
		return model.Auth{}, fmt.Errorf("you've initialized the vault")
	}
	a, _ := as.r.InitMasterPass(hashed) 
	return a, nil
}

func (as *AuthService) Login(masterPass string) (model.Session, error) {
	if !as.r.IsInitialized() {
		return model.Session{}, fmt.Errorf("you haven't initialized the vault")
	}
	hashed, _ := as.r.GetHashedMasterPass()
	res := utils.CompareHashPassword(masterPass, hashed)
	if !res {
		return model.Session{}, fmt.Errorf("the master password is incorrect")
	}
	session := model.Session{
		MasterPassword: masterPass,
		Authenticated:  true,
	}
	return session, nil
}
