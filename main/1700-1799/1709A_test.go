package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1709/problem/A
// https://codeforces.com/problemset/status/1709/problem/A
func TestCF1709A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
0 1 2
1
0 3 2
2
3 1 0
2
1 3 0
outputCopy
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1709A)
}
