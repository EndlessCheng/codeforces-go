package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1491/problem/B
// https://codeforces.com/problemset/status/1491/problem/B
func TestCF1491B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 4
2 2
2 3 4
3 2
2 4 3
3 2
outputCopy
7
3
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1491B)
}
