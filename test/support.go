package test

import (
	"github.com/google/go-cmp/cmp"
	"go/types"
)

func difference(value1 interface{}, value2 interface{}, opts ...cmp.Option) string {
	return cmp.Diff(value1, value2, opts...)
}

func contains[V interface{} | types.Basic](array []V, value V, opts ...cmp.Option) bool {
	for _, v := range array {
		if cmp.Equal(v, value, opts...) {
			return true
		}
	}
	return false
}
