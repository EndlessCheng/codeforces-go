package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1194/E
// https://codeforces.com/problemset/status/1194/problem/E
func TestCF1194E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
-1 4 -1 -2
6 -1 -2 -1
-2 3 6 3
2 -2 2 4
4 -1 4 3
5 3 5 1
5 2 1 2
outputCopy
7
inputCopy
5
1 5 1 0
0 1 5 1
5 4 0 4
4 2 4 0
4 3 4 5
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1194E)
}
