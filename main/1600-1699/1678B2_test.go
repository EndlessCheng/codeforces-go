package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1678/B2
// https://codeforces.com/problemset/status/1678/problem/B2
func TestCF1678B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
10
1110011000
8
11001111
2
00
2
11
6
100110
outputCopy
3 2
0 3
0 1
0 1
3 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1678B2)
}
