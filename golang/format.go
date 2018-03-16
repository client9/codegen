package golang

import (
	"go/format"
)

// Format formats golang codein the cannonical way.
// It is just a wrapper around github.com/golang/go/format@Source
//
// https://github.com/golang/go/blob/master/src/go/format/format.go
//
func Format(src []byte) ([]byte, error) {
	return format.Source(src)
}
