package bykovstorage

import "gitlab.com/bkvstorage/bsencryption"

type bsEncryptor struct {
	cipher  bsencryption.BSCipher
	cryptID string
	//ciphers []bsencryption.BSCipher
}

func (encryptor bsEncryptor) getCypherNames() []string {
	var lCyphers []string
	for _, cypher := range bsencryption.Ciphers {
		lCyphers = append(lCyphers, cypher.GetCipherName())
	}
	return lCyphers
}

func (encryptor *bsEncryptor) Init(cryptID string) error {
	for _, cypher := range bsencryption.Ciphers {
		if cypher.GetGryptID() == cryptID {
			encryptor.cipher = cypher
			encryptor.cryptID = cryptID
			encryptor.cipher.CleanAndInit()
			return nil
		}
	}
	return formError(BSERR00004EncCypherNotExist, "bsEncryptor.Init", "CryptID: "+cryptID)
}

func (encryptor bsEncryptor) getCryptIDbyName(cypherName string) (string, error) {
	for _, cypher := range bsencryption.Ciphers {
		if cypher.GetCipherName() == cypherName {
			return cypher.GetGryptID(), nil
		}
	}
	return "", formError(BSERR00004EncCypherNotExist, "bsEncryptor.getCryptIDbyName", "cypherName: "+cypherName)
}

func (encryptor *bsEncryptor) Encrypt(plainText string) (string, error) {
	encString, err := encryptor.cipher.Encrypt(plainText)
	if err != nil {
		return "", formError(BSERR00008EncEncryptionError, err.Error(), encryptor.cipher.GetCipherName())
	}
	return encString, nil
}

func (encryptor *bsEncryptor) Decrypt(plainText string) (string, error) {
	return encryptor.cipher.Decrypt(plainText)
}

func (encryptor *bsEncryptor) SetPassword(password string) error {
	return encryptor.cipher.SetPassword(password)
}

func (encryptor *bsEncryptor) IsReady() bool {
	return encryptor.cipher.IsKeyGenerated()
}
