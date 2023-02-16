package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1203/problem/E
// https://codeforces.com/problemset/status/1203/problem/E
func TestCF1203E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 2 4 1
outputCopy
4
inputCopy
6
1 1 1 4 4 4
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1203E)
}
