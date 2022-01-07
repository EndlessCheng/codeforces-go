package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1621/problem/E
// https://codeforces.com/problemset/status/1621/problem/E
func TestCF1621E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1
30
3
25 16 37
4 2
9 12 12 6
2
4 5
3
111 11 11
outputCopy
101
00100`
	testutil.AssertEqualCase(t, rawText, 0, CF1621E)
}
