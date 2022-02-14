package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1626/problem/B
// https://codeforces.com/problemset/status/1626/problem/B
func TestCF1626B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
10057
90
outputCopy
10012
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1626B)
}
