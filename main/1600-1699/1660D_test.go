package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1660/problem/D
// https://codeforces.com/problemset/status/1660/problem/D
func TestCF1660D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
1 2 -1 2
3
1 1 -2
5
2 0 -2 2 -1
3
-2 -1 -1
3
-1 -2 -2
outputCopy
0 2
3 0
2 0
0 1
1 0`
	testutil.AssertEqualCase(t, rawText, 0, CF1660D)
}
