package service

import (
	"govault/internal/crypto"
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

// Get secret with decrypted password
func (vs *VaultService) GetSecretById(id string) (model.SecretInVault, error) {
	secret, _ := vs.r.GetSecretById(id)
	pass, _ := crypto.DecryptAES(secret.Ciphertext, secret.DerivedKey)
	res := model.SecretInVault{
		ID:        secret.ID,
		Name:      secret.Name,
		Username:  secret.Username,
		Password:  pass,
		Note:      secret.Note,
		CreatedAt: secret.CreatedAt,
	}
	return res, nil
}
