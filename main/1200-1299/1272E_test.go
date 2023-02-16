package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1272/problem/E
// https://codeforces.com/problemset/status/1272/problem/E
func TestCF1272E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
4 5 7 6 7 5 4 4 6 4
outputCopy
1 1 1 2 -1 1 1 3 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1272E)
}
