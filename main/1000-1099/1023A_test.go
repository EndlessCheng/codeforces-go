package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1023/A
// https://codeforces.com/problemset/status/1023/problem/A
func TestCF1023A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 10
code*s
codeforces
outputCopy
YES
inputCopy
6 5
vk*cup
vkcup
outputCopy
YES
inputCopy
1 1
v
k
outputCopy
NO
inputCopy
9 6
gfgf*gfgf
gfgfgf
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1023A)
}
