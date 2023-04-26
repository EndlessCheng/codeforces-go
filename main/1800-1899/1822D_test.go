package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1822/problem/D
// https://codeforces.com/problemset/status/1822/problem/D
func TestCF1822D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
2
3
6
outputCopy
1
2 1
-1
6 5 2 3 4 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1822D)
}
