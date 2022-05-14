package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1628/B
// https://codeforces.com/problemset/status/1628/problem/B
func TestCF1628B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5
zx
ab
cc
zx
ba
2
ab
bad
4
co
def
orc
es
3
a
b
c
3
ab
cd
cba
2
ab
ab
outputCopy
YES
NO
NO
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1628B)
}
