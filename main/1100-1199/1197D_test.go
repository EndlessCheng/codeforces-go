package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1197/D
// https://codeforces.com/problemset/status/1197/problem/D
func TestCF1197D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3 10
2 -4 15 -3 4 8 3
outputCopy
7
inputCopy
5 2 1000
-13 -4 -9 -20 -11
outputCopy
0
inputCopy
5 10 10
1000000000 1000000000 1000000000 1000000000 1000000000
outputCopy
4999999990`
	testutil.AssertEqualCase(t, rawText, 0, CF1197D)
}
