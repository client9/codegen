package golang

import (
	"bytes"
	"io"
	"strings"
	"text/scanner"
)

func Equal(a, b io.Reader) bool {
	var s1, s2 scanner.Scanner
	s1.Init(a)
	s2.Init(b)
	for {
		t1, t2 := s1.Scan(), s2.Scan()
		if t1 == scanner.EOF && t2 == scanner.EOF {
			return true
		}
		if t1 != t2 || s1.TokenText() != s2.TokenText() {
			return false
		}
		// keep going - tokens are equal and not EOF
	}

}

func EqualString(a, b string) bool {
	if a == b {
		return true
	}
	return Equal(strings.NewReader(a), strings.NewReader(b))
}

// https://github.com/golang/go/blob/master/src/go/format/format.go
func EqualBytes(a, b []byte) bool {
	// fast path, both inputs are indentical
	if bytes.Equal(a, b) {
		return true
	}
	return Equal(bytes.NewReader(a), bytes.NewReader(b))
}
