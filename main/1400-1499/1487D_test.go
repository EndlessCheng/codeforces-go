package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1487/problem/D
// https://codeforces.com/problemset/status/1487/problem/D
func TestCF1487D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
6
9
outputCopy
0
1
1
inputCopy
1
25
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF1487D)
}
