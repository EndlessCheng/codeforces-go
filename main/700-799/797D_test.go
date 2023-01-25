package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/797/problem/D
// https://codeforces.com/problemset/status/797/problem/D
func TestCF797D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
15 -1 -1
10 1 3
5 -1 -1
outputCopy
2
inputCopy
8
6 2 3
3 4 5
12 6 7
1 -1 8
4 -1 -1
5 -1 -1
14 -1 -1
2 -1 -1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF797D)
}
