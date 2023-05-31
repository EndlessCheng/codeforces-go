package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/359/problem/C
// https://codeforces.com/problemset/status/359/problem/C
func TestCF359C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
2 2
outputCopy
8
inputCopy
3 3
1 2 3
outputCopy
27
inputCopy
2 2
29 29
outputCopy
73741817
inputCopy
4 5
0 0 0 0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 2, CF359C)
}
