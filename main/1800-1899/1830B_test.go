package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1830/problem/B
// https://codeforces.com/problemset/status/1830/problem/B
func TestCF1830B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
2 3 2
3 3 1
8
4 2 8 2 1 2 7 5
3 5 8 8 1 1 6 5
8
4 4 8 8 8 8 8 8
8 8 8 8 8 8 8 8
outputCopy
2
7
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1830B)
}
