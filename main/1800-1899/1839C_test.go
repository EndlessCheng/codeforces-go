package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1839/problem/C
// https://codeforces.com/problemset/status/1839/problem/C
func TestCF1839C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 1 0 0 0
1
1
3
0 1 1
6
1 0 0 1 1 0
outputCopy
YES
0 0 2 1 3
NO
NO
YES
0 1 0 2 4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1839C)
}
