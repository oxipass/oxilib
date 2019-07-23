package bykovstorage

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

const dbFile = "test.sqlite"

//const dbFile = ":memory:"

var dbPass = ""

func TestLogin(t *testing.T) {
	// Getting storage instance
	bsInstance := GetInstance()

	errPass := bsInstance.SetPassword(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
}

func TestAddItem(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.SetPassword(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(JSONInputAddItem{ItemName: itemName, ItemIcon: "icon"})
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("Rsponse is not success"))
		t.FailNow()
		return
	}
	bsInstance.Lock()

	errPass = bsInstance.SetPassword(dbPass)
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
	t.Error(errors.New("Created item is not found"))
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
	bsInstance := GetInstance()

	// Setting temporary SQLite DB file
	fullFilename := os.TempDir() + "/bs-" + generateRandomString(4) + dbFile
	fmt.Println(fullFilename)

	errOpen := bsInstance.Open(fullFilename)

	if errOpen != nil {
		return errOpen
	}
	fmt.Println("File open")

	// Generating random password for the database
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
	return os.Remove(dbFile)
}
