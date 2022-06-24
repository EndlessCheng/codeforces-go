package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1283/E
// https://codeforces.com/problemset/status/1283/problem/E
func TestCF1283E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 4 4
outputCopy
2 4
inputCopy
9
1 1 8 8 8 4 4 4 4
outputCopy
3 8
inputCopy
7
4 3 7 1 4 3 3
outputCopy
3 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1283E)
}
