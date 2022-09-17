package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1729/problem/C
// https://codeforces.com/problemset/status/1729/problem/C
func TestCF1729C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
logic
codeforces
bca
aaaaaaaaaaa
adbaadabad
to
outputCopy
9 4
1 4 3 5
16 10
1 8 3 4 9 5 2 6 7 10
1 2
1 3
0 11
1 8 10 4 3 5 7 2 9 6 11
3 10
1 9 5 4 7 3 8 6 2 10
5 2
1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1729C)
}
