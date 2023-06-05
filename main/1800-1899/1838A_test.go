package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1838/problem/A
// https://codeforces.com/problemset/status/1838/problem/A
func TestCF1838A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
3
9 2 7
3
15 -4 11
4
-9 1 11 -10
5
3 0 0 0 3
7
8 16 8 0 8 16 8
4
0 0 0 0
10
27 1 24 28 2 -1 26 25 28 27
6
600000000 800000000 0 -200000000 1000000000 800000000
3
0 -1000000000 1000000000
outputCopy
9
11
-9
3
8
0
-1
600000000
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1838A)
}
