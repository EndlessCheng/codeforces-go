package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/468/A
// https://codeforces.com/problemset/status/468/problem/A
func TestCF468A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
NO
inputCopy
8
outputCopy
YES
8 * 7 = 56
6 * 5 = 30
3 - 4 = -1
1 - 2 = -1
30 - -1 = 31
56 - 31 = 25
25 + -1 = 24
inputCopy
9
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF468A)
}
