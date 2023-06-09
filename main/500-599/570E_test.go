package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/570/E
// https://codeforces.com/problemset/status/570/problem/E
func TestCF570E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
aaab
baaa
abba
outputCopy
3
inputCopy
5 2
ab
ab
cc
ba
ba
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF570E)
}
