package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1770/problem/D
// https://codeforces.com/problemset/status/1770/problem/D
func TestCF1770D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
1 2 2
1 3 3
5
3 3 1 3 4
4 5 2 5 5
outputCopy
6
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1770D)
}
