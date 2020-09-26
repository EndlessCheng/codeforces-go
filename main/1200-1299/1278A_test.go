package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1278A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
abacaba
zyxaabcaabkjh
onetwothree
threetwoone
one
zzonneyy
one
none
twenty
ten
outputCopy
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1278A)
}
