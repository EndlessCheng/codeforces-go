package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1256/E
// https://codeforces.com/problemset/status/1256/problem/E
func TestCF1256E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1 3 4 2
outputCopy
3 1
1 1 1 1 1 
inputCopy
6
1 5 12 13 2 15
outputCopy
7 2
2 2 1 1 2 1 
inputCopy
10
1 2 5 129 185 581 1041 1909 1580 8150
outputCopy
7486 3
3 3 3 2 2 2 2 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1256E)
}
