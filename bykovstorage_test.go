package bykovstorage

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var dbFile = "/tmp/test.sqlite"
var dbPass = ""

func TestLogin(t *testing.T) {
	// Getting storage instance
	bsInstance := GetInstance()

	// Opening file
	/* already should be open
	errOpen := bsInstance.Open(dbFile)
	if errOpen != nil {
		t.Error(errOpen)
		t.FailNow()
		return
	}
	*/

	// Checking password
	errPass := bsInstance.SetPassword(dbPass)
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
	/*
		file, err := ioutil.TempFile(os.TempDir(), "test_file.sqlite")
		if err != nil {
			return err
		}
		dbFile = file.Name()
		fmt.Println("DB file: " + dbFile)
		errClose := file.Close()
		if errClose != nil {
			return errClose
		}
		fmt.Println("File closed")
	*/
	bsInstance := GetInstance()

	errOpen := bsInstance.Open(dbFile)
	if errOpen != nil {
		return errOpen
	}
	fmt.Println("File open")

	dbPass = generateRandomString(12)
	fmt.Println("DB password: " + dbPass)
	errSetPassword := bsInstance.SetNewPassword(dbPass, "AES256")
	if errSetPassword != nil {
		return errSetPassword
	}
	fmt.Println("Password is set")

	return nil
}

func teardown() error {
	if dbFile != "not_set" {
		//return nil
		return os.Remove(dbFile)
	}
	return errors.New("File was not set, do not call teardown without setup")
}
