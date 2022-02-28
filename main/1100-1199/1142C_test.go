package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1142/problem/C
// https://codeforces.com/problemset/status/1142/problem/C
func TestCF1142C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
-1 0
0 2
1 0
outputCopy
2
inputCopy
5
1 0
1 -1
0 -1
-1 0
-1 -1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1142C)
}
