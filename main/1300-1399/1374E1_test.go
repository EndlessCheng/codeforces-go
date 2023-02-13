package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1374/problem/E1
// https://codeforces.com/problemset/status/1374/problem/E1
func TestCF1374E1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 4
7 1 1
2 1 1
4 0 1
8 1 1
1 0 1
1 1 1
1 0 1
3 0 0
outputCopy
18
inputCopy
5 2
6 0 0
9 0 0
1 0 1
2 1 1
5 1 0
outputCopy
8
inputCopy
5 3
3 0 0
2 1 0
3 1 0
5 0 1
3 0 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1374E1)
}
