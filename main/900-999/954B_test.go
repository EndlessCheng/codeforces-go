package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/954/problem/B
// https://codeforces.com/problemset/status/954/problem/B
func TestCF954B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
abcabca
outputCopy
5
inputCopy
8
abcdefgh
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF954B)
}
