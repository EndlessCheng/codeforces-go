package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1111C(t *testing.T) {
	// just copy from website
	rawText := `
2 2 1 2
1 3
outputCopy
6
inputCopy
3 2 1 2
1 7
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, -1, Sol1111C)
}
