package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1036/problem/A
// https://codeforces.com/problemset/status/1036/problem/A
func TestCF1036A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
outputCopy
1
inputCopy
4 12
outputCopy
3
inputCopy
999999999999999999 999999999999999986
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1036A)
}
