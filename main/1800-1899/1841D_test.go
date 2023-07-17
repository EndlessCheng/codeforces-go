package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1841/problem/D
// https://codeforces.com/problemset/status/1841/problem/D
func TestCF1841D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7
2 4
9 12
2 4
7 7
4 8
10 13
6 8
5
2 2
2 8
0 10
1 2
5 6
4
1 1
2 2
3 3
4 4
outputCopy
1
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1841D)
}
