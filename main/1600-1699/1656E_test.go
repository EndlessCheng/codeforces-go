package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1656/problem/E
// https://codeforces.com/problemset/status/1656/problem/E
func TestCF1656E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
1 2
1 3
3 4
3 5
3
1 2
1 3
outputCopy
-3 5 1 2 2
1 1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1656E)
}
