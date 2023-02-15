package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1095/E
// https://codeforces.com/problemset/status/1095/problem/E
func TestCF1095E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
(((())
outputCopy
3
inputCopy
6
()()()
outputCopy
0
inputCopy
1
)
outputCopy
0
inputCopy
8
)))(((((
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1095E)
}
