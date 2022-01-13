package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1016/problem/E
// https://codeforces.com/problemset/status/1016/problem/E
func TestCF1016E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
-3 1 6
2
2 4
6 7
5
3 1
1 3
6 1
6 4
7 6
outputCopy
5.000000000000000
3.000000000000000
0.000000000000000
1.500000000000000
2.000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1016E)
}
