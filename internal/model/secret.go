package model

import "time"

type Secret struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string `gorm:"unique"`
	Username   string
	Ciphertext []byte // AES-encrypted password
	Note       string
	CreatedAt  time.Time
}
