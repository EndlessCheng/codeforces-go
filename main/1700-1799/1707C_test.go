package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1707/C
// https://codeforces.com/problemset/status/1707/problem/C
func TestCF1707C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2
3 5
1 3
3 2
4 2
outputCopy
01111
inputCopy
10 11
1 2
2 5
3 4
4 2
8 1
4 5
10 5
9 5
8 2
5 7
4 6
outputCopy
0011111011`
	testutil.AssertEqualCase(t, rawText, 0, CF1707C)
}
