package test

import (
	"github.com/google/go-cmp/cmp"
	"go/types"
	"strings"
	"testing"
)

//goland:noinspection GoUnusedExportedFunction
func ShouldBeTrue(t *testing.T, condition bool, format string, v ...any) bool {
	t.Helper()
	if condition {
		return true
	}
	t.Error(format, v)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeFalse(t *testing.T, condition bool, format string, v ...any) bool {
	t.Helper()
	if !condition {
		return true
	}
	t.Error(format, v)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if cmp.Equal(value1, value2, opts...) {
		return true
	}
	t.Errorf("Failed: expecting equality:\n"+
		"\tExpected: '%v'\n"+
		"\tGot: '%v'\n\t--- diff ---\n%v",
		value1, value2, difference(value1, value2, opts...))
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if !cmp.Equal(value1, value2, opts...) {
		return true
	}
	t.Errorf("Failed: expecing inequality but found equal '%v'", value1)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeInArray[V interface{} | types.Basic](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	t.Helper()
	if contains(array, value, opts...) {
		return true
	}
	t.Errorf("Value '%v' not found in array", value)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeInArray[V interface{} | string](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	t.Helper()
	if !contains(array, value, opts...) {
		return true
	}
	t.Errorf("Value '%v' found in array", value)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeSubstring(t *testing.T, str string, sub string) bool {
	t.Helper()
	if strings.Contains(str, sub) {
		return true
	}
	t.Errorf("'%s' is not within '%s", sub, str)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeSubString(t *testing.T, str string, sub string) bool {
	t.Helper()
	if !strings.Contains(str, sub) {
		return true
	}
	t.Errorf("'%s' is within '%s", sub, str)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeNoError(t *testing.T, err error, format string, v ...any) bool {
	t.Helper()
	if err != nil {
		t.Errorf(format, v...)
		return false
	}
	return true
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeError(t *testing.T, err error, format string, v ...any) bool {
	t.Helper()
	if err == nil {
		t.Errorf(format, v...)
		return false
	}
	return true
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeNil(t *testing.T, value interface{}, format string, v ...any) bool {
	t.Helper()
	if value != nil {
		return true
	}
	t.Errorf(format, v...)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeNil(t *testing.T, value interface{}, format string, v ...any) bool {
	t.Helper()
	if value == nil {
		return true
	}
	t.Errorf(format, v...)
	return false
}
