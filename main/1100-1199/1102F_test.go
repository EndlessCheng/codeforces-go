package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1102/F
// https://codeforces.com/problemset/status/1102/problem/F
func TestCF1102F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
9 9
10 8
5 3
4 3
outputCopy
5
inputCopy
2 4
1 2 3 4
10 3 7 3
outputCopy
0
inputCopy
6 1
3
6
2
5
1
4
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF1102F)
}
