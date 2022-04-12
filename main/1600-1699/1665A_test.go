package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1665/A
// https://codeforces.com/problemset/status/1665/problem/A
func TestCF1665A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
7
8
9
10
outputCopy
1 1 1 1
2 2 2 1
2 2 2 2
2 4 2 1
3 5 1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1665A)
}
