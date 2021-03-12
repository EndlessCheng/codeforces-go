package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1450/C1
// https://codeforces.com/problemset/status/1450/problem/C1
func TestCF1450C1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
.X.
XXX
.X.
6
XX.XXX
XXXXXX
XXX.XX
XXXXXX
XX.X.X
XXXXXX
5
XXX.X
.X..X
XXX.X
..X..
..X..
outputCopy
.X.
XOX
.X.
XX.XXO
XOXXOX
OXX.XX
XOOXXO
XX.X.X
OXXOXX
XOX.X
.X..X
XXO.O
..X..
..X..`
	testutil.AssertEqualCase(t, rawText, 0, CF1450C1)
}
