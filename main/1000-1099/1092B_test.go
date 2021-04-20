package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1092/problem/B
// https://codeforces.com/problemset/status/1092/problem/B
func TestCF1092B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 10 2 3 14 5
outputCopy
5
inputCopy
2
1 100
outputCopy
99`
	testutil.AssertEqualCase(t, rawText, 0, CF1092B)
}
