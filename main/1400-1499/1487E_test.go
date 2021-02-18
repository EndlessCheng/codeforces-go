package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1487/problem/E
// https://codeforces.com/problemset/status/1487/problem/E
func TestCF1487E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3 2 1
1 2 3 4
5 6 7
8 9
10
2
1 2
1 1
2
3 1
3 2
1
1 1
outputCopy
26
inputCopy
1 1 1 1
1
1
1
1
1
1 1
0
0
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1487E)
}
