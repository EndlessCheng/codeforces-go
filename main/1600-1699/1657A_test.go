package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1657/A
// https://codeforces.com/problemset/status/1657/problem/A
func TestCF1657A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8 6
0 0
9 15
outputCopy
1
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1657A)
}
