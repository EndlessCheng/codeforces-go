package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/835/D
// https://codeforces.com/problemset/status/835/problem/D
func TestCF835D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abba
outputCopy
6 1 0 0 
inputCopy
abacaba
outputCopy
12 4 1 0 0 0 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF835D)
}
