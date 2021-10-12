package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1177/B
// https://codeforces.com/problemset/status/1177/problem/B
func TestCF1177B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
outputCopy
7
inputCopy
21
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1177B)
}
