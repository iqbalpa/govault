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
func (vs *VaultService) GetSecretById(masterPass, id string) (model.SecretInVault, error) {
	secret, _ := vs.r.GetSecretById(id)
	pass, _ := crypto.DecryptAES(masterPass, secret.Ciphertext, secret.Salt)
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

// Create a new secret
func (vs *VaultService) CreateSecret(masterPass, name, username, password, note string, salt []byte) (model.Secret, error) {
	ciphertext, _ := crypto.EncryptAES(masterPass, password, salt)
	secret, _ := vs.r.CreateSecret(name, username, note, ciphertext, salt)
	return secret, nil
}

// Delete a secret
func (vs *VaultService) DeleteSecretById(id string) (model.Secret, error) {
	secret, _ := vs.r.DeleteSecretById(id)
	return secret, nil
}
