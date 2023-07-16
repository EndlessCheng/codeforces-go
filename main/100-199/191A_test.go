package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/191/A
// https://codeforces.com/problemset/status/191/problem/A
func TestCF191A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
abc
ca
cba
outputCopy
6
inputCopy
4
vvp
vvp
dam
vvp
outputCopy
0
inputCopy
3
ab
c
def
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF191A)
}
