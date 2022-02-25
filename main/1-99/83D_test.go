package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/83/D
// https://codeforces.com/problemset/status/83/problem/D
func TestCF83D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 10 2
outputCopy
5
inputCopy
12 23 3
outputCopy
2
inputCopy
6 19 5
outputCopy
0
inputCopy
1 2000000000 44711
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF83D)
}
