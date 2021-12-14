package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1401/B
// https://codeforces.com/problemset/status/1401/problem/B
func TestCF1401B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 2
3 3 1
4 0 1
2 3 0
0 0 1
0 0 1
outputCopy
4
2
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1401B)
}
