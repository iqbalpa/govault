package model

import "time"

type Secret struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string `gorm:"unique"`
	Username   string
	Ciphertext []byte // AES-encrypted password
	DerivedKey []byte // Derived key from master password
	Note       string
	CreatedAt  time.Time
}

type SecretInVault struct {
	ID        string
	Name      string
	Username  string
	Password  string
	Note      string
	CreatedAt time.Time
}
