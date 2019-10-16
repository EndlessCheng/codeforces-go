package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol567D(t *testing.T) {
	// just copy from website
	rawText := `
11 3 3
5
4 8 6 1 11
outputCopy
3
inputCopy
5 1 3
2
1 5
outputCopy
-1
inputCopy
5 1 3
1
3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, Sol567D)
}
