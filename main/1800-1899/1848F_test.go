package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1848/problem/F
// https://codeforces.com/problemset/status/1848/problem/F
func TestCF1848F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 1 2
outputCopy
2
inputCopy
2
0 0
outputCopy
0
inputCopy
1
14
outputCopy
1
inputCopy
8
0 1 2 3 4 5 6 7
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1848F)
}
