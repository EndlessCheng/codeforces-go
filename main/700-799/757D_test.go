package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/757/D
// https://codeforces.com/problemset/status/757/problem/D
func TestCF757D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1011
outputCopy
10
inputCopy
2
10
outputCopy
1
inputCopy
62
00010011000110010011110110011001110110010011110110111100100010
outputCopy
`
	testutil.AssertEqualCase(t, rawText, -1, CF757D)
}
