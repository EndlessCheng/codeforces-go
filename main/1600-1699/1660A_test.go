package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1660/problem/A
// https://codeforces.com/problemset/status/1660/problem/A
func TestCF1660A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
4 0
0 2
0 0
2314 2374
outputCopy
4
5
1
1
7063`
	testutil.AssertEqualCase(t, rawText, 0, CF1660A)
}
