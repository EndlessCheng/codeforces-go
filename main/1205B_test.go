package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1205B(t *testing.T) {
	// just copy from website
	rawText := `
4
3 6 28 9
outputCopy
4
inputCopy
5
5 12 9 16 48
outputCopy
3
inputCopy
4
1 2 4 8
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, -1, Sol1205B)
}
