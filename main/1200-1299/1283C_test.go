package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1283/C
// https://codeforces.com/problemset/status/1283/problem/C
func TestCF1283C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 0 0 2 4
outputCopy
5 3 1 2 4 
inputCopy
7
7 0 0 1 4 0 6
outputCopy
7 3 2 1 4 5 6 
inputCopy
7
7 4 0 3 0 5 1
outputCopy
7 4 2 3 6 5 1 
inputCopy
5
2 1 0 0 0
outputCopy
2 1 4 5 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1283C)
}
