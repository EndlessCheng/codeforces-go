package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/862/E
// https://codeforces.com/problemset/status/862/problem/E
func TestCF862E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6 3
1 2 3 4 5
1 2 3 4 5 6
1 1 10
1 1 -9
1 5 -1
outputCopy
0
9
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF862E)
}
