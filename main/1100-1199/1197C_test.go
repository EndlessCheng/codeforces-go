package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1197/problem/C
// https://codeforces.com/problemset/status/1197/problem/C
func TestCF1197C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
4 8 15 16 23 42
outputCopy
12
inputCopy
4 4
1 3 3 7
outputCopy
0
inputCopy
8 1
1 1 2 3 5 8 13 21
outputCopy
20`
	testutil.AssertEqualCase(t, rawText, 0, CF1197C)
}
