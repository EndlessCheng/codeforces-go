package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1775/problem/E
// https://codeforces.com/problemset/status/1775/problem/E
func TestCF1775E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
1 2 -3
5
1 0 0 -1 -1
6
2 -4 3 -5 4 1
5
1 -1 1 -1 1
7
0 0 0 0 0 0 0
outputCopy
3
2
6
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1775E)
}
