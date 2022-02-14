package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1606/problem/C
// https://codeforces.com/problemset/status/1606/problem/C
func TestCF1606C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 13
0 1 2
2 777
0 4
3 255
0 1 3
10 1000000000
0 1 2 3 4 5 6 7 8 9
outputCopy
59
778
148999
999999920999999999`
	testutil.AssertEqualCase(t, rawText, 1, CF1606C)
}
