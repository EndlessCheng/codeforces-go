package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1644/problem/B
// https://codeforces.com/problemset/status/1644/problem/B
func TestCF1644B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
3
outputCopy
4 1 3 2
1 2 4 3
3 4 1 2
2 4 1 3
3 2 1
1 3 2
3 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1644B)
}
