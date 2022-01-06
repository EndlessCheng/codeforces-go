package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1620/C
// https://codeforces.com/problemset/status/1620/problem/C
func TestCF1620C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 4 3
a*
4 1 3
a**a
6 3 20
**a***
outputCopy
abb
abba
babbbbbbbbb`
	testutil.AssertEqualCase(t, rawText, 0, CF1620C)
}
