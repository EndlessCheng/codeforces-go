package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/932/C
// https://codeforces.com/problemset/status/932/problem/C
func TestCF932C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 2 5
outputCopy
6 5 8 3 4 1 9 2 7
inputCopy
3 2 1
outputCopy
1 2 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF932C)
}
