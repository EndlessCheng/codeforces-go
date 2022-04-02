package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1660/problem/E
// https://codeforces.com/problemset/status/1660/problem/E
func TestCF1660E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4

3
010
011
100

5
00010
00001
10000
01000
00100

2
10
10

4
1111
1011
1111
1111
outputCopy
1
0
2
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1660E)
}
