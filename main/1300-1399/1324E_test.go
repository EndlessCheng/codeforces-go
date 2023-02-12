package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1324/problem/E
// https://codeforces.com/problemset/status/1324/problem/E
func TestCF1324E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 24 21 23
16 17 14 20 20 11 22
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1324E)
}
