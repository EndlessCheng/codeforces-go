package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1557/problem/D
// https://codeforces.com/problemset/status/1557/problem/D
func TestCF1557D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 6
1 1 1
1 7 8
2 7 7
2 15 15
3 1 1
3 15 15
outputCopy
0
inputCopy
5 4
1 2 3
2 4 6
3 3 5
5 1 1
outputCopy
3
2 4 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1557D)
}
