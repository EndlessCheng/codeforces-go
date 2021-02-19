package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1486/problem/E
// https://codeforces.com/problemset/status/1486/problem/E
func TestCF1486E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
1 2 3
2 3 4
3 4 5
4 5 6
1 5 1
2 4 2
outputCopy
0 98 49 25 114 
inputCopy
3 2
1 2 1
2 3 2
outputCopy
0 -1 9 `
	testutil.AssertEqualCase(t, rawText, 0, CF1486E)
}
