package security

import (
	"github.com/oxipass/oxicrypt"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
)

func (enc OxiEncryptor) GetCypherNames() []string {
	var lCyphers []string
	for _, cypher := range oxicrypt.Ciphers {
		lCyphers = append(lCyphers, cypher.GetCipherName())
	}
	return lCyphers
}

func (enc *OxiEncryptor) Init(cryptID string) error {
	for _, cypher := range oxicrypt.Ciphers {
		if cypher.GetCryptID() == cryptID {
			enc.Cipher = cypher
			enc.CryptID = cryptID
			enc.Cipher.CleanAndInit()
			return nil
		}
	}
	return oxierr.FormError(oxierr.BSERR00004EncCypherNotExist, "OxiEncryptor.Init", "CryptID: "+cryptID)
}

func (enc OxiEncryptor) GetCryptIDbyName(cypherName string) (string, error) {
	for _, cypher := range oxicrypt.Ciphers {
		if cypher.GetCipherName() == cypherName {
			return cypher.GetCryptID(), nil
		}
	}
	return "", oxierr.FormError(oxierr.BSERR00004EncCypherNotExist, "OxiEncryptor.getCryptIDbyName", "cypherName: "+cypherName)
}

func (enc *OxiEncryptor) Encrypt(plainText string) (string, error) {
	if enc.Cipher == nil {
		return "", oxierr.FormError(oxierr.BSERR00008EncEncryptionError, "encryptor is not initialized")
	}
	encString, err := enc.Cipher.Encrypt(plainText)
	if err != nil {
		return "", oxierr.FormError(oxierr.BSERR00008EncEncryptionError, err.Error(), enc.Cipher.GetCipherName())
	}
	return encString, nil
}

func (enc *OxiEncryptor) Decrypt(plainText string) (string, error) {
	return enc.Cipher.Decrypt(plainText)
}

func (enc *OxiEncryptor) SetPassword(password string) error {
	return enc.Cipher.SetPassword(password)
}

func (enc *OxiEncryptor) IsReady() bool {
	return enc.Cipher.IsPasswordSet()
}
