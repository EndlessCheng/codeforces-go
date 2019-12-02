package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol107B(t *testing.T) {
	// just copy from website
	rawText := `
3 2 1
2 1
outputCopy
1
inputCopy
3 2 1
1 1
outputCopy
-1
inputCopy
3 2 1
2 2
outputCopy
0.666667`
	testutil.AssertEqualCase(t, rawText, 0, Sol107B)
}
