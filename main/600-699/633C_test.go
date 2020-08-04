package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF633C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
30
ariksihsidlihcdnaehsetahgnisol
10
Kira
hates
is
he
losing
death
childish
L
and
Note
outputCopy
Kira is childish and he hates losing 
inputCopy
12
iherehtolleh
5
HI
Ho
there
HeLLo
hello
outputCopy
HI there HeLLo `
	testutil.AssertEqualCase(t, rawText, 0, CF633C)
}
