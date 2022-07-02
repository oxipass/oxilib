package oxilib

import "github.com/oxipass/oxicrypt"

type bsEncryptor struct {
	cipher  oxicrypt.BSCipher
	cryptID string
}

func (enc bsEncryptor) getCypherNames() []string {
	var lCyphers []string
	for _, cypher := range oxicrypt.Ciphers {
		lCyphers = append(lCyphers, cypher.GetCipherName())
	}
	return lCyphers
}

func (enc *bsEncryptor) Init(cryptID string) error {
	for _, cypher := range oxicrypt.Ciphers {
		if cypher.GetCryptID() == cryptID {
			enc.cipher = cypher
			enc.cryptID = cryptID
			enc.cipher.CleanAndInit()
			return nil
		}
	}
	return formError(BSERR00004EncCypherNotExist, "bsEncryptor.Init", "CryptID: "+cryptID)
}

func (enc bsEncryptor) getCryptIDbyName(cypherName string) (string, error) {
	for _, cypher := range oxicrypt.Ciphers {
		if cypher.GetCipherName() == cypherName {
			return cypher.GetCryptID(), nil
		}
	}
	return "", formError(BSERR00004EncCypherNotExist, "bsEncryptor.getCryptIDbyName", "cypherName: "+cypherName)
}

func (enc *bsEncryptor) Encrypt(plainText string) (string, error) {
	if enc.cipher == nil {
		return "", formError(BSERR00008EncEncryptionError, "encryptor is not initialized")
	}
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
	return enc.cipher.IsPasswordSet()
}
