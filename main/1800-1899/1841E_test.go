package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1841/problem/E
// https://codeforces.com/problemset/status/1841/problem/E
func TestCF1841E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3
0 0 0
9
4
2 0 3 1
5
4
2 0 3 1
6
4
2 0 3 1
10
10
0 2 2 1 5 10 3 4 1 1
20
1
1
0
outputCopy
6
3
4
4
16
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1841E)
}
