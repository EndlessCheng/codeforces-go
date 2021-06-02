package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/165/E
// https://codeforces.com/problemset/status/165/problem/E
func TestCF165E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
90 36
outputCopy
36 90
inputCopy
4
3 6 3 6
outputCopy
-1 -1 -1 -1
inputCopy
5
10 6 9 8 2
outputCopy
-1 8 2 2 8`
	testutil.AssertEqualCase(t, rawText, 0, CF165E)
}
