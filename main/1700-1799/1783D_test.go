package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1783/D
// https://codeforces.com/problemset/status/1783/problem/D
func TestCF1783D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1 1
outputCopy
3
inputCopy
5
1 2 3 5 0
outputCopy
7
inputCopy
18
74 92 52 36 48 51 33 64 43 95 92 78 35 77 55 25 50 6
outputCopy
64960`
	testutil.AssertEqualCase(t, rawText, 0, CF1783D)
}
