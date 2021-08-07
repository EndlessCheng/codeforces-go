package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1065/B
// https://codeforces.com/problemset/status/1065/problem/B
func TestCF1065B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
outputCopy
0 1
inputCopy
3 1
outputCopy
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1065B)
}
