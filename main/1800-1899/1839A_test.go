package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1839/problem/A
// https://codeforces.com/problemset/status/1839/problem/A
func TestCF1839A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 2
5 2
9 3
7 1
10 4
9 5
8 8
outputCopy
2
3
4
7
4
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1839A)
}
