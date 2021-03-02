package bslib

import (
	"strings"
	"testing"
)

const nonExistingCypher = "blah-blah"

func TestInitNonExistingCypher(t *testing.T) {
	newCypher := new(bsEncryptor)
	err := newCypher.Init(nonExistingCypher)
	if err == nil {
		t.Error("Should return error, cypher does not exist")
	} else if !strings.Contains(err.Error(), BSERR00004EncCypherNotExist) {
		t.Error("Wrong error message, cypher does not exist")
	}
}

func TestCypherNames(t *testing.T) {
	newCypher := new(bsEncryptor)
	names := newCypher.getCypherNames()
	if names == nil || len(names) <= 0 {
		t.Error("Cannot retrieve cypher's names")
	}

}

func TestEncryptWithoutInit(t *testing.T) {
	newCypher := new(bsEncryptor)
	_, err := newCypher.Encrypt(nonExistingCypher)
	if err == nil {
		t.Error("Should return error, cypher is not initialized")
	}
}

const cCryptId = "8GA63DMN" // AES-256
func TestEncryptEmptyString(t *testing.T) {
	newCypher := new(bsEncryptor)
	err := newCypher.Init(cCryptId)
	if err != nil {
		t.Error(err.Error())
		return
	}
	_, err = newCypher.Encrypt("")
	if err == nil {
		t.Error("Should return error, cypher is not initialized")
	}
}

func TestCheckNonExistingCypher(t *testing.T) {
	newCypher := new(bsEncryptor)
	_, err := newCypher.getCryptIDbyName(nonExistingCypher)
	if err == nil {
		t.Error("Should return error, cypher is fake")
	}
}
