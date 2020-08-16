package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1209D(t *testing.T) {
	// just copy from website
	rawText := `
5 4
1 2
4 3
1 4
3 4
outputCopy
1
inputCopy
6 5
2 3
2 1
3 4
6 5
4 5
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, Sol1209D)
}
