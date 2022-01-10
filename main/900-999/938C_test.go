package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/938/problem/C
// https://codeforces.com/problemset/status/938/problem/C
func TestCF938C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
21
0
1
outputCopy
5 2
1 1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF938C)
}
