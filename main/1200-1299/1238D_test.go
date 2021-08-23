package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1238/D
// https://codeforces.com/problemset/status/1238/problem/D
func TestCF1238D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
AABBB
outputCopy
6
inputCopy
3
AAA
outputCopy
3
inputCopy
7
AAABABB
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1238D)
}
