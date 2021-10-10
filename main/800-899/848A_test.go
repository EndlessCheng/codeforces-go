package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/848/A
// https://codeforces.com/problemset/status/848/problem/A
func TestCF848A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12
outputCopy
abababab
inputCopy
3
outputCopy
codeforces
inputCopy
0
outputCopy
a`
	testutil.AssertEqualCase(t, rawText, 0, CF848A)
}
