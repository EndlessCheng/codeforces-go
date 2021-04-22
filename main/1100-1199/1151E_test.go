package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1151/E
// https://codeforces.com/problemset/status/1151/problem/E
func TestCF1151E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 1 3
outputCopy
7
inputCopy
4
2 1 1 3
outputCopy
11
inputCopy
10
1 5 2 5 5 3 10 6 5 1
outputCopy
104`
	testutil.AssertEqualCase(t, rawText, 0, CF1151E)
}
