package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1530/E
// https://codeforces.com/problemset/status/1530/problem/E
func TestCF1530E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
vkcup
abababa
zzzzzz
outputCopy
ckpuv
aababab
zzzzzz
inputCopy
3
aaaaaaaaaabbb
aaaaaaaabbccc
aaabbbccc
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1530E)
}
