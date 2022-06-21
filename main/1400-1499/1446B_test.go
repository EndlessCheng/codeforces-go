package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1446/B
// https://codeforces.com/problemset/status/1446/problem/B
func TestCF1446B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
abba
babab
outputCopy
5
inputCopy
8 10
bbbbabab
bbbabaaaaa
outputCopy
12
inputCopy
7 7
uiibwws
qhtkxcn
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1446B)
}
