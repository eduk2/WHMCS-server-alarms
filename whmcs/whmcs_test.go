package whmcs

import (
	"testing"
)

func TestGetDatabaseConnection(t *testing.T) {
	configurationFilePath := "configurationExample.php"

	dbHost, dbUser, dbPass, dbName := getDatabaseConnection(configurationFilePath)

	varsListToTest := []struct {
		searchingName  string
		searchingValue string
		expected       string
	}{
		{"dbHost", dbHost, "localhost"},
		{"dbUser", dbUser, "user1"},
		{"dbPass", dbPass, "pass1"},
		{"dbName", dbName, "name1"},
	}

	for _, varOfList := range varsListToTest {
		if varOfList.searchingValue == "" {
			t.Fatalf("Error: %v is empty", varOfList.searchingName)
		}

		if varOfList.searchingValue != varOfList.expected {
			t.Fatalf("Error: %v has not been obtained successfully:\nExpected = %v\nGot = %v", varOfList.searchingName, varOfList.expected, varOfList.searchingValue)
		}
	}

}
