package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1163/D
// https://codeforces.com/problemset/status/1163/problem/D
func TestCF1163D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
*****
katie
shiro
outputCopy
1
inputCopy
caat
caat
a
outputCopy
-1
inputCopy
*a*
bba
b
outputCopy
0
inputCopy
***
cc
z
outputCopy
2
inputCopy
*********
abc
bca
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF1163D)
}
