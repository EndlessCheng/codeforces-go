package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1626/problem/E
// https://codeforces.com/problemset/status/1626/problem/E
func TestCF1626E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
0 1 0 0 0 0 1 0
8 6
2 5
7 8
6 5
4 5
6 1
7 3
outputCopy
0 1 1 1 1 0 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1626E)
}
