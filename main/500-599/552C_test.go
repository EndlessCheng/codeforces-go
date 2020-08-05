package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol552C(t *testing.T) {
	// just copy from website
	rawText := `
4 11
outputCopy
YES
inputCopy
4 7
outputCopy
NO
inputCopy
3 7
outputCopy
YES
inputCopy
100 99
outputCopy
YES
inputCopy
100 50
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, Sol552C)
}
