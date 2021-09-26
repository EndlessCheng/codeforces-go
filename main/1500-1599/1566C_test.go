package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1566/C
// https://codeforces.com/problemset/status/1566/problem/C
func TestCF1566C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7
0101000
1101100
5
01100
10101
2
01
01
6
000000
111111
outputCopy
8
8
2
12`
	testutil.AssertEqualCase(t, rawText, 0, CF1566C)
}
