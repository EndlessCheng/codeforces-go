package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/460/D
// https://codeforces.com/problemset/status/460/problem/D
func TestCF460D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 15 3
outputCopy
1
2
10 11
inputCopy
8 30 7
outputCopy
0
5
14 9 28 11 16
inputCopy
262143 393216 3
outputCopy
0
3
262143 393215 393216`
	testutil.AssertEqualCase(t, rawText, -1, CF460D)
}
