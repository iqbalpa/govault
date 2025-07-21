package service

import (
	"encoding/json"
	"fmt"
	"govault/internal/crypto"
	"govault/internal/model"
	"govault/internal/repository"
	"os"
	"path"
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
	secrets, err := vs.r.GetAllSecrets()
	if err != nil {
		return []model.Secret{}, err
	}
	return secrets, nil
}

// Get secret with decrypted password
func (vs *VaultService) GetSecretById(masterPass, id string) (model.SecretInVault, error) {
	secret, err := vs.r.GetSecretById(id)
	if err != nil {
		return model.SecretInVault{}, err
	}
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
func (vs *VaultService) CreateSecret(masterPass, name, username, password, note string, salt []byte) (model.SecretInVault, error) {
	ciphertext, _ := crypto.EncryptAES(masterPass, password, salt)
	secret, err := vs.r.CreateSecret(name, username, note, ciphertext, salt)
	if err != nil {
		return model.SecretInVault{}, err
	}
	res := model.SecretInVault{
		ID:        secret.ID,
		Name:      secret.Name,
		Username:  secret.Username,
		Note:      secret.Note,
		CreatedAt: secret.CreatedAt,
	}
	return res, nil
}

// Delete a secret
func (vs *VaultService) DeleteSecretById(id string) (model.SecretInVault, error) {
	secret, err := vs.r.DeleteSecretById(id)
	if err != nil {
		return model.SecretInVault{}, err
	}
	res := model.SecretInVault{
		ID:        secret.ID,
		Name:      secret.Name,
		Username:  secret.Username,
		Note:      secret.Note,
		CreatedAt: secret.CreatedAt,
	}
	return res, nil
}

// Export secrets to JSON
func (vs *VaultService) ExportAllSecrets(fpath string) error {
	secrets, err := vs.r.GetAllSecrets()
	if err != nil {
		return err
	}

	j, _ := json.Marshal(secrets)
	absPath, _ := os.Getwd()
	f, _ := os.Create(path.Join(absPath, fpath))
	_, err = f.Write(j)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
