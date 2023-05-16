package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1827/C
// https://codeforces.com/problemset/status/1827/problem/C
func TestCF1827C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6
abaaba
1
a
2
aa
6
abcdef
12
accabccbacca
6
abbaaa
outputCopy
3
0
1
0
14
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1827C)
}
