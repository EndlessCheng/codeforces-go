package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1649/problem/A
// https://codeforces.com/problemset/status/1649/problem/A
func TestCF1649A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
1 1
5
1 0 1 0 1
4
1 0 1 1
outputCopy
0
4
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1649A)
}
