package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1882/problem/D
// https://codeforces.com/problemset/status/1882/problem/D
func TestCF1882D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
3 2 1 0
1 2
2 3
2 4
1
100
outputCopy
8 6 12 10 
0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1882D)
}
