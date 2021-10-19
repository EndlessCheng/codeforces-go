package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/682/C
// https://codeforces.com/problemset/status/682/problem/C
func TestCF682C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
88 22 83 14 95 91 98 53 11
3 24
7 -8
1 67
1 64
9 65
5 12
6 -80
3 8
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 1, CF682C)
}
