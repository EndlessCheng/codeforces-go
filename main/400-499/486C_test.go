package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol486C(t *testing.T) {
	// just copy from website
	rawText := `
4 1
abba
outputCopy
0
inputCopy
4 1
abbz
outputCopy
1
inputCopy
8 3
aeabcaez
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, -1, Sol486C)
}
