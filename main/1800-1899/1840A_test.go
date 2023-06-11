package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1840/problem/A
// https://codeforces.com/problemset/status/1840/problem/A
func TestCF1840A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8
abacabac
5
qzxcq
20
ccooddeeffoorrcceess
outputCopy
ac
q
codeforces`
	testutil.AssertEqualCase(t, rawText, 0, CF1840A)
}
