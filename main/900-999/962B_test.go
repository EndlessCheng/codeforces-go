package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/962/B
// https://codeforces.com/problemset/status/962/problem/B
func TestCF962B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 1
*...*
outputCopy
2
inputCopy
6 2 3
*...*.
outputCopy
4
inputCopy
11 3 10
.*....**.*.
outputCopy
7
inputCopy
3 2 3
***
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF962B)
}
