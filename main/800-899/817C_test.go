package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/817/C
// https://codeforces.com/problemset/status/817/problem/C
func TestCF817C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12 1
outputCopy
3
inputCopy
25 20
outputCopy
0
inputCopy
10 9
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF817C)
}
