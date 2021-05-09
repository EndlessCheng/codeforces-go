package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/471/D
// https://codeforces.com/problemset/status/471/problem/D
func TestCF471D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
13 5
2 4 5 5 4 3 2 2 2 3 3 2 1
3 4 4 3 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF471D)
}
