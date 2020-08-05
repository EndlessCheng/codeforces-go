package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol439D(t *testing.T) {
	// just copy from website
	rawText := `
2 2
2 3
3 5
outputCopy
3
inputCopy
3 2
1 2 3
3 4
outputCopy
4
inputCopy
3 2
4 5 6
1 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, Sol439D)
}
