package test

import (
	"github.com/google/go-cmp/cmp"
	"go/types"
	"strings"
	"testing"
)

var fatalErrors = true

//goland:noinspection GoUnusedExportedFunction
func SetErrorsFatal(cond bool) {
	fatalErrors = cond
}

func failed(t *testing.T, format string, v ...any) bool {
	if fatalErrors {
		t.Fatalf(format, v...)
	}
	t.Errorf(format, v...)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeTrue(t *testing.T, condition bool, format string, v ...any) bool {
	t.Helper()
	if condition {
		return true
	}
	return failed(t, format, v...)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeFalse(t *testing.T, condition bool, format string, v ...any) bool {
	t.Helper()
	if !condition {
		return true
	}
	return failed(t, format, v...)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if cmp.Equal(value1, value2, opts...) {
		return true
	}
	return failed(t, "Failed: expecting equality:\n"+
		"\tExpected: '%v'\n"+
		"\tGot: '%v'\n\t--- diff ---\n%v",
		value1, value2, difference(value1, value2, opts...))
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if !cmp.Equal(value1, value2, opts...) {
		return true
	}
	return failed(t, "Failed: expecting inequality but found equal '%v'", value1)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeInArray[V interface{} | types.Basic](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	t.Helper()
	if contains(array, value, opts...) {
		return true
	}
	return failed(t, "Value '%v' not found in array", value)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeInArray[V interface{} | string](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	t.Helper()
	if !contains(array, value, opts...) {
		return true
	}
	return failed(t, "Value '%v' found in array", value)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeSubstring(t *testing.T, str string, sub string) bool {
	t.Helper()
	if strings.Contains(str, sub) {
		return true
	}
	return failed(t, "'%s' is not within '%s", sub, str)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeSubString(t *testing.T, str string, sub string) bool {
	t.Helper()
	if !strings.Contains(str, sub) {
		return true
	}
	return failed(t, "'%s' is within '%s", sub, str)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeNoError(t *testing.T, err error, format string, v ...any) bool {
	t.Helper()
	if err == nil {
		return true
	}
	return failed(t, format, v...)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeError(t *testing.T, err error, format string, v ...any) bool {
	t.Helper()
	if err != nil {
		return true
	}
	return failed(t, format, v...)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeNil(t *testing.T, value interface{}, format string, v ...any) bool {
	t.Helper()
	if value != nil {
		return true
	}
	return failed(t, format, v...)
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeNil(t *testing.T, value interface{}, format string, v ...any) bool {
	t.Helper()
	if value == nil {
		return true
	}
	return failed(t, format, v...)
}
