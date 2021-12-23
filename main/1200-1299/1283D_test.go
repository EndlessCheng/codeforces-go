package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1283/D
// https://codeforces.com/problemset/status/1283/problem/D
func TestCF1283D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 6
1 5
outputCopy
8
-1 2 6 4 0 3 
inputCopy
3 5
0 3 1
outputCopy
7
5 -2 4 -1 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1283D)
}
