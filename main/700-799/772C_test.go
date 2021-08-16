package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/772/C
// https://codeforces.com/problemset/status/772/problem/C
func TestCF772C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 5
outputCopy
5
1 2 4 3 0
inputCopy
3 10
2 9 1
outputCopy
6
3 9 2 9 8 0`
	testutil.AssertEqualCase(t, rawText, 0, CF772C)
}
