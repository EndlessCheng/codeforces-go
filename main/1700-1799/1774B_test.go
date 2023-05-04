package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1774/problem/B
// https://codeforces.com/problemset/status/1774/problem/B
func TestCF1774B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
12 6 2
1 1 1 1 1 7
12 6 2
2 2 2 2 2 2
outputCopy
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1774B)
}
