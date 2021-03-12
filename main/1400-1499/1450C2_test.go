package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1450/C2
// https://codeforces.com/problemset/status/1450/problem/C2
func TestCF1450C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
.O.
OOO
.O.
6
XXXOOO
XXXOOO
XX..OO
OO..XX
OOOXXX
OOOXXX
5
.OOO.
OXXXO
OXXXO
OXXXO
.OOO.
outputCopy
.O.
OXO
.O.
OXXOOX
XOXOXO
XX..OO
OO..XX
OXOXOX
XOOXXO
.OXO.
OOXXO
XXOXX
OXXOO
.OXO.`
	testutil.AssertEqualCase(t, rawText, 0, CF1450C2)
}
