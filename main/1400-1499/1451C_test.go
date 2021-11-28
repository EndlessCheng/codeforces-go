package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1451/C
// https://codeforces.com/problemset/status/1451/problem/C
func TestCF1451C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 3
abc
bcd
4 2
abba
azza
2 1
zz
aa
6 2
aaabba
ddddcc
outputCopy
No
Yes
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1451C)
}
