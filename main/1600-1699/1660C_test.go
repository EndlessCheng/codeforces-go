package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1660/problem/C
// https://codeforces.com/problemset/status/1660/problem/C
func TestCF1660C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
aabbdabdccc
zyx
aaababbb
aabbcc
oaoaaaoo
bmefbmuyw
outputCopy
3
3
2
0
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1660C)
}
