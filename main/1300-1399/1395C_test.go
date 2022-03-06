package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1395/C
// https://codeforces.com/problemset/status/1395/problem/C
func TestCF1395C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
2 6 4 0
2 4
outputCopy
2
inputCopy
7 6
1 9 1 9 8 1 0
1 1 4 5 1 4
outputCopy
0
inputCopy
8 5
179 261 432 162 82 43 10 38
379 357 202 184 197
outputCopy
147`
	testutil.AssertEqualCase(t, rawText, 0, CF1395C)
}
