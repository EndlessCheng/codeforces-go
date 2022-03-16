package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1006/problem/E
// https://codeforces.com/problemset/status/1006/problem/E
func TestCF1006E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 6
1 1 1 3 5 3 5 7
3 1
1 5
3 4
7 3
1 8
1 9
outputCopy
3
6
8
-1
9
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1006E)
}
