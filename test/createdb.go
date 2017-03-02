package test

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/errors"
)

func init() {
	for _, suite := range []string{SuitePouchLocal, SuitePouchRemote, SuiteCouch16, SuiteCouch20, SuiteKivikMemory, SuiteCloudant,
		SuitePouchRemoteNoAuth, SuiteCouch16NoAuth, SuiteCouch20NoAuth, SuiteCloudantNoAuth} { //FIXME: SuiteKivikServer,SuiteKivikServerNoAuth
		RegisterTest(suite, "CreateDB", true, CreateDB)
	}
}

// CreateDB tests database creation.
func CreateDB(client *kivik.Client, suite string, fail FailFunc) {
	testDB := testDBName()
	fmt.Printf("testDB = %s\n", testDB)
	// defer client.DestroyDB(testDB)
	err := client.CreateDB(testDB)
	if strings.Contains(suite, "NoAuth") {
		switch errors.StatusCode(err) {
		case 0:
			fail("CreateDB: Should fail for unauthenticated session")
		case http.StatusUnauthorized:
			// Expected
		default:
			fail("CreateDB: Expected 401/Unauthorized, Got: %s", err)
		}
		return
	}
	if err != nil {
		fail("Failed to create database '%s': %s", testDB, err)
	}
}
