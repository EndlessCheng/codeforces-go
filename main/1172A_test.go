package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1172A(t *testing.T) {
	// just copy from website
	rawText := `
3
0 2 0
3 0 1
outputCopy
2
inputCopy
3
0 2 0
1 0 3
outputCopy
4
inputCopy
11
0 0 0 5 0 0 0 4 0 0 11
9 2 6 0 8 1 7 0 3 0 10
outputCopy
18`
	testutil.AssertEqualCase(t, rawText, 0, Sol1172A)
}
