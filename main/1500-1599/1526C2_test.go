package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1526/C2
// https://codeforces.com/problemset/status/1526/problem/C2
func TestCF1526C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 -4 1 -3 1 -3
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1526C2)
}
