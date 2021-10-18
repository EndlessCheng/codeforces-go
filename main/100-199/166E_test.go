package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/166/E
// https://codeforces.com/problemset/status/166/problem/E
func TestCF166E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
3
inputCopy
4
outputCopy
21`
	testutil.AssertEqualCase(t, rawText, 0, CF166E)
}
