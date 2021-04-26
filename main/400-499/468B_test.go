package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/468/problem/B
// https://codeforces.com/problemset/status/468/problem/B
func TestCF468B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5 9
2 3 4 5
outputCopy
YES
0 0 1 1
inputCopy
3 3 4
1 2 4
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF468B)
}
