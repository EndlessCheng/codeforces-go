package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1296/C
// https://codeforces.com/problemset/status/1296/problem/C
func TestCF1296C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
LRUD
4
LURD
5
RRUDU
5
LLDDR
outputCopy
1 2
1 4
3 4
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1296C)
}
