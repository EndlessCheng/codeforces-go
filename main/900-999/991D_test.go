package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/991/D
// https://codeforces.com/problemset/status/991/problem/D
func TestCF991D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
00
00
outputCopy
1
inputCopy
00X00X0XXX0
0XXX0X00X00
outputCopy
4
inputCopy
0X0X0
0X0X0
outputCopy
0
inputCopy
0XXX0
00000
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF991D)
}
