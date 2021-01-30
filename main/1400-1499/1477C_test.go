package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1477/C
// https://codeforces.com/problemset/status/1477/problem/C
func TestCF1477C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 0
5 0
4 2
2 1
3 0
outputCopy
1 2 5 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1477C)
}
