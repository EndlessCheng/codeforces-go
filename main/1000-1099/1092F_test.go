package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1092/problem/F
// https://codeforces.com/problemset/status/1092/problem/F
func TestCF1092F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
9 4 1 7 10 1 6 5
1 2
2 3
1 4
1 5
5 6
5 7
5 8
outputCopy
121
inputCopy
1
1337
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 1, CF1092F)
}
