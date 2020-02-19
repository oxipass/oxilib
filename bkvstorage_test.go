package bslib

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

var fullPathDBFile string
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

func TestDeleteItem(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(JSONInputUpdateItem{ItemName: itemName, ItemIcon: "fas fa-ambulance"})
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful"))
		t.FailNow()
		return
	}
	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	delResponse, errDelete := bsInstance.DeleteItem(JSONInputUpdateItem{ItemID: response.ItemID})
	if errDelete != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	if delResponse.Status != ConstSuccessResponse {
		t.Error(errors.New("deletion response is not successful"))
		t.FailNow()
		return
	}

	errLock = bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	items, respErr := bsInstance.ReadAllItems()
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	for _, item := range items {
		if item.ID == response.ItemID && item.Deleted == false {
			t.Error(errors.New("deleted item is found as not deleted"))
			t.FailNow()
			return
		}
	}

}

func TestAddItemWithNonExistingIcon(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(JSONInputUpdateItem{ItemName: itemName, ItemIcon: "ewfdwejfnerfkj"})
	if err == nil {
		t.Error(errors.New("no error returned in spite of the fact that icon is not existing"))
		t.FailNow()
		return
	}
	if response.Status == ConstSuccessResponse {
		t.Error(errors.New("item was added in spite of the fact that icon is not existing"))
		t.FailNow()
		return
	}
}

func TestAddItem(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(JSONInputUpdateItem{ItemName: itemName, ItemIcon: "fas fa-ambulance"})
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successfull"))
		t.FailNow()
		return
	}
	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	items, respErr := bsInstance.ReadAllItems()
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}
	for _, item := range items {
		if item.ID == response.ItemID && item.Name == itemName {
			return
		}
	}
	t.Error(errors.New("created item is not found"))
	t.FailNow()

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
		return errors.New("cannot retrieve BSLib instance")
	}
	log.Println("BSLib instance successfully retrieved")

	// Setting temporary SQLite DB file
	// fullPathDBFile = generateTempFilename()
	// FIXME: change hardcoded back to temporary file name
	fullPathDBFile = "/Users/bykov/etc/bsweb/bs.sqlite"
	log.Println("Full path to database file: " + fullPathDBFile)

	if _, err := os.Stat(fullPathDBFile); err == nil {
		err = os.Remove(fullPathDBFile)
		if err != nil {
			return err
		}
	}

	errOpen := bsInstance.Open(fullPathDBFile)
	if errOpen != nil {
		return errOpen
	}
	log.Println("BSLib is initiated successfully")

	// Generating random password for the database
	dbPass = generateRandomString(12)
	log.Println("Generated DB password: " + dbPass)

	errSetPassword := bsInstance.SetNewPassword(dbPass, "AES256V1")
	if errSetPassword != nil {
		return errSetPassword
	}
	log.Println("Password is successfully set")

	errLock := bsInstance.Lock()
	if errLock != nil {
		return errLock
	}
	log.Println("Database is locked, the tests can be started")

	return nil
}

func teardown() error {
	if _, err := os.Stat(fullPathDBFile); err == nil {
		// FIXME: uncomment file removal before submitting to git
		//err = os.Remove(fullPathDBFile)
		if err != nil {
			return err
		}
	}

	return nil
}
