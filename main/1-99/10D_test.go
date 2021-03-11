package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/10/problem/D
// https://codeforces.com/problemset/status/10/problem/D
func TestCF10D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
2 3 1 6 5 4 6
4
1 3 5 6
outputCopy
3
3 5 6 
inputCopy
5
1 2 0 2 1
3
1 0 1
outputCopy
2
0 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF10D)
}
