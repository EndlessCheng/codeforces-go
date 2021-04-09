package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1285/problem/D
// https://codeforces.com/problemset/status/1285/problem/D
func TestCF1285D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
outputCopy
2
inputCopy
2
1 5
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1285D)
}
