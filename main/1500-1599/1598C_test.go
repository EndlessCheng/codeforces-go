package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1598/problem/C
// https://codeforces.com/problemset/status/1598/problem/C
func TestCF1598C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
8 8 8 8
3
50 20 10
5
1 4 7 3 5
7
1 2 3 4 5 6 7
outputCopy
6
0
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1598C)
}
