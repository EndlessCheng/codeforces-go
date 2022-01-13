package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1117/problem/C
// https://codeforces.com/problemset/status/1117/problem/C
func TestCF1117C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 0
4 6
3
UUU
outputCopy
5
inputCopy
0 3
0 0
3
UDD
outputCopy
3
inputCopy
0 0
0 1
1
L
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1117C)
}
