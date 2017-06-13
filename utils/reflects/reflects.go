//
// Package reflects contains helpers to the run-time reflection, allowing
// a program to manipulate objects with arbitrary types
//
package reflects

import (
	"path"
	"reflect"
	"runtime"
)

// GetFuncName returns the name of the function.
func GetFuncName(i interface{}) string {
	return path.Ext(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name())
}
