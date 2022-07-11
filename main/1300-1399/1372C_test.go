package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1372/C
// https://codeforces.com/problemset/status/1372/problem/C
func TestCF1372C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
1 2 3 4 5
7
3 2 4 5 1 6 7
outputCopy
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1372C)
}
