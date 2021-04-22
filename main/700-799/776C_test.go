package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/776/C
// https://codeforces.com/problemset/status/776/problem/C
func TestCF776C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
2 2 2 2
outputCopy
8
inputCopy
4 -3
3 -6 -3 12
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF776C)
}
