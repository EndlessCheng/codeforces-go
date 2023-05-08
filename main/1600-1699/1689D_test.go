package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1689/D
// https://codeforces.com/problemset/status/1689/problem/D
func TestCF1689D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2
BW
WW
WB
3 3
WWB
WBW
BWW
2 3
BBB
BBB
5 5
BWBWB
WBWBW
BWBWB
WBWBW
BWBWB
9 9
WWWWWWWWW
WWWWWWWWW
BWWWWWWWW
WWWWWWWWW
WWWWBWWWW
WWWWWWWWW
WWWWWWWWW
WWWWWWWWW
WWWWWWWWB
outputCopy
2 1
2 2
1 2
3 3
6 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1689D)
}
