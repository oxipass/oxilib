package models

import "github.com/oxipass/oxicrypt"

type OxiEncryptor struct {
	cipher  oxicrypt.BSCipher
	cryptID string
}
