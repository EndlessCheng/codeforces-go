package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1621/problem/D
// https://codeforces.com/problemset/status/1621/problem/D
func TestCF1621D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
0 8
1 99
2
0 0 0 0
0 0 0 0
9 9 2 2
9 9 9 9
2
0 0 4 2
0 0 2 4
4 2 4 2
2 4 2 4
4
0 0 0 0 0 0 0 2
0 0 0 0 0 0 2 0
0 0 0 0 0 2 0 0
0 0 0 0 2 0 0 0
0 0 0 2 2 0 2 2
0 0 2 0 1 6 2 1
0 2 0 0 2 4 7 4
2 0 0 0 2 0 1 6
outputCopy
100
22
14
42`
	testutil.AssertEqualCase(t, rawText, 0, CF1621D)
}
