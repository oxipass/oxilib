package bykovstorage

import (
	"github.com/bykovme/bsencrypt"
)

type bsEncryptor struct {
	cipher  bsencrypt.BSCipher
	cryptID string
}

func (enc bsEncryptor) getCypherNames() []string {
	var lCyphers []string
	for _, cypher := range bsencrypt.Ciphers {
		lCyphers = append(lCyphers, cypher.GetCipherName())
	}
	return lCyphers
}

func (enc *bsEncryptor) Init(cryptID string) error {
	for _, cypher := range bsencrypt.Ciphers {
		if cypher.GetGryptID() == cryptID {
			enc.cipher = cypher
			enc.cryptID = cryptID
			enc.cipher.CleanAndInit()
			return nil
		}
	}
	return formError(BSERR00004EncCypherNotExist, "bsEncryptor.Init", "CryptID: "+cryptID)
}

func (enc bsEncryptor) getCryptIDbyName(cypherName string) (string, error) {
	for _, cypher := range bsencrypt.Ciphers {
		if cypher.GetCipherName() == cypherName {
			return cypher.GetGryptID(), nil
		}
	}
	return "", formError(BSERR00004EncCypherNotExist, "bsEncryptor.getCryptIDbyName", "cypherName: "+cypherName)
}

func (enc *bsEncryptor) Encrypt(plainText string) (string, error) {
	encString, err := enc.cipher.Encrypt(plainText)
	if err != nil {
		return "", formError(BSERR00008EncEncryptionError, err.Error(), enc.cipher.GetCipherName())
	}
	return encString, nil
}

func (enc *bsEncryptor) Decrypt(plainText string) (string, error) {
	return enc.cipher.Decrypt(plainText)
}

func (enc *bsEncryptor) SetPassword(password string) error {
	return enc.cipher.SetPassword(password)
}

func (enc *bsEncryptor) IsReady() bool {
	return enc.cipher.IsKeyGenerated()
}
