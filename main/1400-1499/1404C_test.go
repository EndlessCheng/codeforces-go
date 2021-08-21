package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1404/C
// https://codeforces.com/problemset/status/1404/problem/C
func TestCF1404C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
13 5
2 2 3 9 5 4 6 5 7 8 3 11 13
3 1
0 0
2 4
5 0
0 12
outputCopy
5
11
6
1
0
inputCopy
5 2
1 4 1 2 4
0 0
1 0
outputCopy
2
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1404C)
}
