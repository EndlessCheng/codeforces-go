package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1677/A
// https://codeforces.com/problemset/status/1677/problem/A
func TestCF1677A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
5 3 6 1 4 2
4
1 2 3 4
10
5 1 6 2 8 3 4 10 9 7
outputCopy
3
0
28`
	testutil.AssertEqualCase(t, rawText, 0, CF1677A)
}
