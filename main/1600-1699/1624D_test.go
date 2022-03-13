package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1624/problem/D
// https://codeforces.com/problemset/status/1624/problem/D
func TestCF1624D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
8 2
bxyaxzay
6 3
aaaaaa
6 1
abcdef
6 6
abcdef
3 2
dxd
11 2
abcabcabcac
6 6
sipkic
7 2
eatoohd
3 1
llw
6 2
bfvfbv
outputCopy
3
2
1
1
1
5
1
1
3
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1624D)
}
