package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1579/problem/E2
// https://codeforces.com/problemset/status/1579/problem/E2
func TestCF1579E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
3 7 5 5
3
3 2 1
3
3 1 2
4
-1 2 2 -1
4
4 5 1 3
5
1 3 1 3 2
outputCopy
2
0
1
0
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1579E2)
}
