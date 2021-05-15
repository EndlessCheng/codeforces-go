package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/gymProblem/101234/G
func Test_runG(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 30
4 2 1 16 8
outputCopy
30
inputCopy
4 5
1 1 2 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, runG)
}