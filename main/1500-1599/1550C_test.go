package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1550/problem/C
// https://codeforces.com/problemset/status/1550/problem/C
func TestCF1550C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
2 4 1 3
5
6 9 1 9 6
2
13 37
outputCopy
10
12
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1550C)
}
