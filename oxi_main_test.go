package oxilib

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

var globalDBFile string
var dbPass = ""

func TestLogin(t *testing.T) {
	// Getting storage instance
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
}

func TestMain(m *testing.M) {
	errSetup := setup()

	if errSetup != nil {
		fmt.Printf("Setup failure: %s\n", errSetup.Error())
		os.Exit(1)
		return
	}

	retCode := m.Run()

	errFinish := teardown()
	if errFinish != nil {
		os.Exit(2)
		return
	}

	os.Exit(retCode)
}

func setup() error {
	// Getting package instance
	log.Println("Tests initial setup")

	bsInstance := GetInstance()
	if bsInstance == nil {
		return errors.New("cannot retrieve oxilib instance")
	}
	log.Println("oxilib instance successfully retrieved")

	dbFile := getTestDbFileName()
	log.Println("Full path to database file: " + dbFile)

	if _, err := os.Stat(dbFile); err == nil {
		err = os.Remove(dbFile)
		if err != nil {
			return err
		}
	}

	errOpen := bsInstance.Open(dbFile)
	if errOpen != nil {
		return errOpen
	}
	log.Println("oxilib is initiated successfully")

	// Generating random password for the database
	dbPass = generateRandomString(12)
	log.Println("Generated DB password: " + dbPass)

	errSetPassword := bsInstance.SetNewPassword(dbPass, "AES256V1")
	if errSetPassword != nil {
		return errSetPassword
	}
	log.Println("Password is successfully set")

	errFillEmpty := bsInstance.FillEmptyStorage()
	if errFillEmpty != nil {
		return errFillEmpty
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		return errLock
	}
	log.Println("Database is locked, the tests can be started")

	return nil
}

func teardown() error {
	bsInstance := GetInstance()
	err := bsInstance.Close()
	if err != nil {
		return nil
	}
	if _, err := os.Stat(getTestDbFileName()); err == nil {
		// Do not delete if local file is used (use build tag 'local' to configure)
		// Check config_global.go for default configuration
		if !isLocalFileUsed() {
			err = os.Remove(getTestDbFileName())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isLocalFileUsed() bool {
	return getTestDbFileName() == localTestFile
}

func getTestDbFileName() string {
	if globalDBFile != "" {
		return globalDBFile
	}

	if useLocalTestFile {
		globalDBFile = localTestFile
		return globalDBFile
	}

	globalDBFile = generateTempFilename()

	return globalDBFile
}
