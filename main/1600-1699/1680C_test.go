package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1680/problem/C
// https://codeforces.com/problemset/status/1680/problem/C
func TestCF1680C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
101110110
1001001001001
0000111111
00000
1111
outputCopy
1
3
0
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1680C)
}
