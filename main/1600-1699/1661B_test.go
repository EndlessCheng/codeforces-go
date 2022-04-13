package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1661/problem/B
// https://codeforces.com/problemset/status/1661/problem/B
func TestCF1661B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
19 32764 10240 49
outputCopy
14 4 4 15 `
	testutil.AssertEqualCase(t, rawText, 0, CF1661B)
}
