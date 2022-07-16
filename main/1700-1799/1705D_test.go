package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1705/D
// https://codeforces.com/problemset/status/1705/problem/D
func TestCF1705D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
0100
0010
4
1010
0100
5
01001
00011
6
000101
010011
outputCopy
2
-1
-1
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1705D)
}
