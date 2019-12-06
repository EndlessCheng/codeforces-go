package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1265D(t *testing.T) {
	// just copy from website
	rawText := `
2 2 2 1
outputCopy
YES
0 1 0 1 2 3 2
inputCopy
1 2 3 4
outputCopy
NO
inputCopy
2 2 2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, Sol1265D)
}
