package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1256/D
// https://codeforces.com/problemset/status/1256/problem/D
func TestCF1256D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8 5
11011010
7 9
1111100
7 11
1111100
outputCopy
01011110
0101111
0011111`
	testutil.AssertEqualCase(t, rawText, 0, CF1256D)
}
