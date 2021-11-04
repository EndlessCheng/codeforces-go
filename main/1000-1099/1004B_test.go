package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1004/B
// https://codeforces.com/problemset/status/1004/problem/B
func TestCF1004B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 3
2 4
2 5
outputCopy
01100
inputCopy
6 3
5 6
1 4
4 6
outputCopy
110010`
	testutil.AssertEqualCase(t, rawText, 0, CF1004B)
}
