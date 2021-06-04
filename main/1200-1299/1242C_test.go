package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1242/problem/C
// https://codeforces.com/problemset/status/1242/problem/C
func TestCF1242C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 1 7 4
2 3 2
2 8 5
1 10
outputCopy
Yes
7 2
2 3
5 1
10 4
inputCopy
2
2 3 -2
2 -1 5
outputCopy
No
inputCopy
2
2 -10 10
2 0 -20
outputCopy
Yes
-10 2
-20 1`
	testutil.AssertEqualCase(t, rawText, -1, CF1242C)
}
