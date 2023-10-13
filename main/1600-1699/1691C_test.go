package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1691/C
// https://codeforces.com/problemset/status/1691/problem/C
func TestCF1691C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 0
1010
7 1
0010100
5 2
00110
outputCopy
21
22
12`
	testutil.AssertEqualCase(t, rawText, 0, CF1691C)
}
