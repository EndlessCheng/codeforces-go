package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/A
// https://codeforces.com/problemset/status/1037/problem/A
func TestCF1037A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
outputCopy
3
inputCopy
2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1037A)
}
