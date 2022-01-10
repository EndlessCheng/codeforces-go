package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1036/problem/B
// https://codeforces.com/problemset/status/1036/problem/B
func TestCF1036B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2 3
4 3 7
10 1 9
outputCopy
1
6
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1036B)
}
