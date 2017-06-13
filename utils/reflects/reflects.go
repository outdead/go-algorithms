//
// Package reflects contains helpers to the run-time reflection, allowing
// a program to manipulate objects with arbitrary types
//
package reflects

import (
	"path"
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName returns the name of the function.
// TODO: Fix panic if i not func
func GetFuncName(i interface{}) string {
	return strings.Trim(path.Ext(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
}
