package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1239/C
// https://codeforces.com/problemset/status/1239/problem/C
func TestCF1239C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 314
0 310 942 628 0
outputCopy
314 628 1256 942 1570 
inputCopy
3 5
5 1 0
outputCopy
15 10 5
inputCopy
3 6
0 5 3
outputCopy
6 12 18
inputCopy
4 5
6 1 4 0
outputCopy
15 10 20 5`
	testutil.AssertEqualCase(t, rawText, -1, CF1239C)
}
