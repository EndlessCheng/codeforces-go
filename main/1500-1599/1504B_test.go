package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1504/B
// https://codeforces.com/problemset/status/1504/problem/B
func TestCF1504B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
10
0111010000
0100101100
4
0000
0000
3
001
000
12
010101010101
100110011010
6
000111
110100
outputCopy
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1504B)
}
