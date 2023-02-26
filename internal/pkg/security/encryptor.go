package security

import "github.com/oxipass/oxicrypt"

type OxiEncryptor struct {
	Cipher  oxicrypt.OxiCipher
	CryptID string
}
