package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1650/problem/D
// https://codeforces.com/problemset/status/1650/problem/D
func TestCF1650D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
3 2 5 6 1 4
3
3 1 2
8
5 8 1 3 2 6 4 7
outputCopy
0 1 1 2 0 4 
0 0 1 
0 1 2 0 2 5 6 2 `
	testutil.AssertEqualCase(t, rawText, 1, CF1650D)
}
