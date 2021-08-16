package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1536/E
// https://codeforces.com/problemset/status/1536/problem/E
func TestCF1536E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 4
0000
00#0
0000
2 1
#
#
1 2
##
6 29
#############################
#000##0###0##0#0####0####000#
#0#0##00#00##00####0#0###0#0#
#0#0##0#0#0##00###00000##00##
#000##0###0##0#0##0###0##0#0#
#############################
outputCopy
2
3
3
319908071`
	testutil.AssertEqualCase(t, rawText, 0, CF1536E)
}
