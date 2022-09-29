package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1178/E
// https://codeforces.com/problemset/status/1178/problem/E
func TestCF1178E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
cacbac
outputCopy
aba
inputCopy
abc
outputCopy
a
inputCopy
cbacacacbcbababacbcb
outputCopy
cbaaacbcaaabc`
	testutil.AssertEqualCase(t, rawText, 0, CF1178E)
}
