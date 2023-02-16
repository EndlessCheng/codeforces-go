package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1579/problem/A
// https://codeforces.com/problemset/status/1579/problem/A
func TestCF1579A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
ABACAB
ABBA
AC
ABC
CABCBB
BCBCBCBCBCBCBCBC
outputCopy
NO
YES
NO
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1579A)
}
