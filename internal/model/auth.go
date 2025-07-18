package model

// I don't need ID, because the master password is only one
type Auth struct {
	HashedMasterPassword string
}
