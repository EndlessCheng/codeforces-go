package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1580/A
// https://codeforces.com/problemset/status/1580/problem/A
func TestCF1580A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
5 4
1000
0000
0110
0000
0001
outputCopy
12
inputCopy
1
9 9
001010001
101110100
000010011
100000001
101010101
110001111
000001111
111100000
000110000
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 1, CF1580A)
}
