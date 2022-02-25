package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1644/problem/C
// https://codeforces.com/problemset/status/1644/problem/C
func TestCF1644C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 2
4 1 3 2
3 5
-2 -7 -1
10 2
-6 -1 -2 4 -6 -1 -4 4 -5 -4
outputCopy
10 12 14 16 18
0 4 4 5
4 6 6 7 7 7 7 8 8 8 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1644C)
}
