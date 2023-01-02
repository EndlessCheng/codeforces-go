package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/379/D
// https://codeforces.com/problemset/status/379/problem/D
func TestCF379D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 2 2
outputCopy
AC
AC
inputCopy
3 3 2 2
outputCopy
Happy new year!
inputCopy
3 0 2 2
outputCopy
AA
AA
inputCopy
4 3 2 1
outputCopy
Happy new year!
inputCopy
4 2 2 1
outputCopy
Happy new year!`
	testutil.AssertEqualCase(t, rawText, 0, CF379D)
}
