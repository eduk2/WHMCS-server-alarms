package whmcs

import (
	v "github.com/eduk2/govarsfromfiles"
)

// get database connection from whmcs
func getDatabaseConnection(configurationFilePath string) (dbHost, dbUser, dbPass, dbName string) {

	if v.File.SetPath(configurationFilePath) {

		v.File.SetRegularExp(`'(.*)'`)
		myValues := v.File.GetValues("$db_host", "$db_username", "$db_name", "$db_password")
		dbHost = myValues["$db_host"]
		dbUser = myValues["$db_username"]
		dbName = myValues["$db_name"]
		dbPass = myValues["$db_password"]

	}

	return
}

// connect to whmcs

// get list of servers
