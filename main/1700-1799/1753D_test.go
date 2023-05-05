package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1753/problem/D
// https://codeforces.com/problemset/status/1753/problem/D
func TestCF1753D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 5
5 2
.LR##
##LR.
outputCopy
4
inputCopy
2 3
4 5
LR.
#.#
outputCopy
-1
inputCopy
4 3
10 10
.LR
###
UU#
DD.
outputCopy
-1
inputCopy
3 6
10 7
.U##.#
#DLR##
.##LR.
outputCopy
24`
	testutil.AssertEqualCase(t, rawText, 0, CF1753D)
}
