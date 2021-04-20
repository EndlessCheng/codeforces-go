package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1092/problem/A
// https://codeforces.com/problemset/status/1092/problem/A
func TestCF1092A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7 3
4 4
6 2
outputCopy
cbcacab
abcd
baabab`
	testutil.AssertEqualCase(t, rawText, 0, CF1092A)
}
