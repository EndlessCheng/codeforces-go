package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1788/D
// https://codeforces.com/problemset/status/1788/problem/D
func TestCF1788D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 4 6
outputCopy
11
inputCopy
5
1 3 5 11 15
outputCopy
30`
	testutil.AssertEqualCase(t, rawText, 0, CF1788D)
}
