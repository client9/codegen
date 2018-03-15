package golang

import (
	"testing"
)

func TestEqual(t *testing.T) {
	var messy = `
package main


import (
   "fmt"
   "strings"
	"text/scanner"
)

func main() {
	const src = "hello world"
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	for tok := s.Scan();
	    tok != scanner.EOF; /*
	    multi
	    line
	    comment
	    */
	    tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}

// comment

`

	var clean = `
package main

import (
        "fmt"
        "strings"
        "text/scanner"
)       

func main() {
        const src = "hello world"
        var s scanner.Scanner
        s.Init(strings.NewReader(src))
        for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
                fmt.Printf("%s: %s\n", s.Position, s.TokenText())
        }   
}`

	var other = `
package main

import (
        "fmt"
        "strings"
        "text/scanner"
)       

func main() {
	const src = "foo bar"
        var s scanner.Scanner
        s.Init(strings.NewReader(src))
        for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
                fmt.Printf("%s: %s\n", s.Position, s.TokenText())
        }   
}`

	var short = `package main`

	if !EqualString(messy, clean) {
		t.Errorf("messy and clean not equals")
	}
	if !EqualString(messy, messy) {
		t.Errorf("messy and messy not equals")
	}
	if !EqualString(clean, clean) {
		t.Errorf("clean and clean not equals")
	}
	if EqualString(messy, other) {
		t.Errorf("messy and other are equal")
	}
	if EqualString(clean, other) {
		t.Errorf("clean and other are equal")
	}
	if EqualString(clean, short) {
		t.Errorf("clean and short are equal")
	}
}
