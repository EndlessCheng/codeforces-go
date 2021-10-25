package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/277/B
// https://codeforces.com/problemset/status/277/problem/B
func TestCF277B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
outputCopy
0 0
3 0
0 3
1 1
inputCopy
6 3
outputCopy
-1
inputCopy
6 6
outputCopy
10 0
-10 0
10 1
9 1
9 -1
0 -2
inputCopy
7 4
outputCopy
176166 6377
709276 539564
654734 174109
910147 434207
790497 366519
606663 21061
859328 886001`
	testutil.AssertEqualCase(t, rawText, 0, CF277B)
}
