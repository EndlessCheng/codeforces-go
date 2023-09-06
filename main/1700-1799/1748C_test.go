package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1748/C
// https://codeforces.com/problemset/status/1748/problem/C
func TestCF1748C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
2 0 1 -1 0
3
1000000000 1000000000 0
4
0 0 0 0
8
3 0 2 -10 10 -30 30 0
9
1 0 0 1 -1 0 1 0 -1
outputCopy
3
1
4
4
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1748C)
}
