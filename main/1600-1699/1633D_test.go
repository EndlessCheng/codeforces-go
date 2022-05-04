package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1633/problem/D
// https://codeforces.com/problemset/status/1633/problem/D
func TestCF1633D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 4
1 7 5 2
2 6 5 2
3 0
3 5 2
5 4 7
5 9
5 2 5 6 3
5 9 1 9 7
6 14
11 4 6 2 8 16
43 45 9 41 15 38
outputCopy
9
0
30
167`
	testutil.AssertEqualCase(t, rawText, 0, CF1633D)
}
