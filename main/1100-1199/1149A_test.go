package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1149/problem/A
// https://codeforces.com/problemset/status/1149/problem/A
func TestCF1149A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 1 2 1
outputCopy
1 1 1 2 2
inputCopy
9
1 1 2 1 1 1 2 1 1
outputCopy
1 1 1 2 1 1 1 2 1
inputCopy
2
1 2
outputCopy
2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1149A)
}
