package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1744/problem/D
// https://codeforces.com/problemset/status/1744/problem/D
func TestCF1744D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1
2
2
3 2
3
10 6 11
4
13 17 1 1
5
1 1 12 1 1
6
20 7 14 18 3 5
outputCopy
0
1
1
-1
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1744D)
}
