package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/25/A
// https://codeforces.com/problemset/status/25/problem/A
func TestCF25A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 4 7 8 10
outputCopy
3
inputCopy
4
1 2 1 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF25A)
}
