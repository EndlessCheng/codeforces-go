package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1237D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
11 5 2 7
outputCopy
1 1 3 2
inputCopy
4
3 2 5 3
outputCopy
5 4 3 6
inputCopy
3
4 3 6
outputCopy
-1 -1 -1`
	testutil.AssertEqualCase(t, rawText, 0, CF1237D)
}
