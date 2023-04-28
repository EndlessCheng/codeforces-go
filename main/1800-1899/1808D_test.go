package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1808/D
// https://codeforces.com/problemset/status/1808/problem/D
func TestCF1808D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 5
1 2 8 2 5 2 8 6
outputCopy
4
inputCopy
9 9
1 2 3 4 5 4 3 2 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1808D)
}
