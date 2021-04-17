package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1509/C
// https://codeforces.com/problemset/status/1509/problem/C
func TestCF1509C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 1 2
outputCopy
3
inputCopy
1
5
outputCopy
0
inputCopy
6
1 6 3 3 6 3
outputCopy
11
inputCopy
6
104 943872923 6589 889921234 1000000000 69
outputCopy
2833800505`
	testutil.AssertEqualCase(t, rawText, 0, CF1509C)
}
