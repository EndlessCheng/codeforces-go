package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1729/G
// https://codeforces.com/problemset/status/1729/problem/G
func TestCF1729G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
abababacababa
aba
ddddddd
dddd
xyzxyz
xyz
abc
abcd
abacaba
abaca
abc
def
aaaaaaaa
a
aaaaaaaa
aa
outputCopy
2 2
1 4
2 1
0 1
1 1
0 1
8 1
3 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1729G)
}
