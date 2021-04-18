package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1301/B
// https://codeforces.com/problemset/status/1301/problem/B
func TestCF1301B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5
-1 10 -1 12 -1
5
-1 40 35 -1 35
6
-1 -1 9 -1 3 -1
2
-1 -1
2
0 -1
4
1 -1 3 -1
7
1 -1 7 5 2 -1 5
outputCopy
1 11
5 35
3 6
0 42
0 0
1 2
3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1301B)
}
