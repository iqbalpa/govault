package service

import (
	"govault/internal/model"
	"govault/internal/repository"
)

type VaultService struct {
	r repository.SecretRepository
}

func New(r repository.SecretRepository) *VaultService {
	return &VaultService{
		r: r,
	}
}

// List all available secrets (without decryption)
func (vs *VaultService) GetAllSecrets() ([]model.Secret, error) {
	secrets, _ := vs.r.GetAllSecrets()
	return secrets, nil
}
