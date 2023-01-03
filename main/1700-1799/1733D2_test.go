package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1733/D2
// https://codeforces.com/problemset/status/1733/problem/D2
func TestCF1733D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 8 9
01001
00101
6 2 11
000001
100000
5 7 2
01000
11011
7 8 3
0111001
0100001
6 3 4
010001
101000
5 10 1
01100
01100
outputCopy
8
10
-1
6
7
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1733D2)
}
