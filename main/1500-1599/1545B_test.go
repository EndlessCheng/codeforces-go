package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1545/B
// https://codeforces.com/problemset/status/1545/problem/B
func TestCF1545B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
0110
6
011011
5
01010
20
10001111110110111000
20
00110110100110111101
20
11101111011000100010
outputCopy
3
6
1
1287
1287
715`
	testutil.AssertEqualCase(t, rawText, 0, CF1545B)
}
