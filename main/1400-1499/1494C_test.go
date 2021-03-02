package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1494/C
// https://codeforces.com/problemset/status/1494/problem/C
func TestCF1494C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 6
-1 1 5 11 15
-4 -3 -2 6 7 15
2 2
-1 1
-1000000000 1000000000
2 2
-1000000000 1000000000
-1 1
3 5
-1 1 2
-2 -1 1 2 5
2 1
1 2
10
outputCopy
4
2
0
3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1494C)
}
