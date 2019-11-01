package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol552D(t *testing.T) {
	// just copy from website
	rawText := `
4
0 0
1 1
2 0
2 2
outputCopy
3
inputCopy
3
0 0
1 1
2 0
outputCopy
1
inputCopy
1
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, Sol552D)
}
