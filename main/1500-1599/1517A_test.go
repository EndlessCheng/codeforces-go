package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1517/problem/A
// https://codeforces.com/problemset/status/1517/problem/A
func TestCF1517A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
205
2050
4100
20500
22550
25308639900
outputCopy
-1
1
2
1
2
36`
	testutil.AssertEqualCase(t, rawText, 0, CF1517A)
}
