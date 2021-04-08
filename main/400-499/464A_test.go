package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/464/problem/A
// https://codeforces.com/problemset/status/464/problem/A
func TestCF464A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
cba
outputCopy
NO
inputCopy
3 4
cba
outputCopy
cbd
inputCopy
4 4
abcd
outputCopy
abda`
	testutil.AssertEqualCase(t, rawText, 0, CF464A)
}
