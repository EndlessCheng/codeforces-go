package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1415/C
// https://codeforces.com/problemset/status/1415/problem/C
func TestCF1415C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
10 3 2
0101010101
2 2
5 4 1
00000
2 10
11 2 3
10110011000
4 3
outputCopy
2
4
10`
	testutil.AssertEqualCase(t, rawText, -1, CF1415C)
}
