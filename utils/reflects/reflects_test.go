//
// Package reflects contains helpers to the run-time reflection, allowing
// a program to manipulate objects with arbitrary types
//
package reflects

import (
	"sort"
	"testing"
)

func testGetFuncName(t *testing.T, checkFunc interface{}, wantName string) {
	if checkName := GetFuncName(checkFunc); checkName != wantName {
		t.Fatalf("GetFuncName returns %s, want, %s", checkName, wantName)
	}
}

func TestGetFuncName(t *testing.T) {
	testGetFuncName(t, sort.Search, "Search")
	testGetFuncName(t, testGetFuncName, "testGetFuncName")
	// testGetFuncName(t, 123, "testGetFuncName") // panic
}
