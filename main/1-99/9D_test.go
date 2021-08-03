package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/9/D
// https://codeforces.com/problemset/status/9/problem/D
func TestCF9D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
outputCopy
5
inputCopy
3 3
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF9D)
}
