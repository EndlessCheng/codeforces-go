package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1239/A
// https://codeforces.com/problemset/status/1239/problem/A
func TestCF1239A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1239A)
}
