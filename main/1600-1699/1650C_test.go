package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1650/problem/C
// https://codeforces.com/problemset/status/1650/problem/C
func TestCF1650C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3

3 8
0 10
-2 1
4 10
11 20
7 -1
9 1
2 3
5 -2

3 6
-1 2
1 3
3 -1
2 4
4 0
8 2

2 5
5 -1
3 -2
1 0
-2 0
-5 -3
outputCopy
12
2 6
5 1
7 8

10
1 6
5 2
3 4

-6
5 1
4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1650C)
}
