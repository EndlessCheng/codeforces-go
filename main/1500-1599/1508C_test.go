package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1508/problem/C
// https://codeforces.com/problemset/status/1508/problem/C
func TestCF1508C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
2 1 14
1 4 14
3 2 15
4 3 8
outputCopy
15
inputCopy
6 6
3 6 4
2 4 1
4 5 7
3 4 10
3 5 1
5 2 15
outputCopy
0
inputCopy
5 6
2 3 11
5 3 7
1 4 10
2 4 14
4 3 8
2 5 6
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1508C)
}
