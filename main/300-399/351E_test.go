package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/351/E
// https://codeforces.com/problemset/status/351/problem/E
func TestCF351E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 1
outputCopy
0
inputCopy
9
-2 0 -1 0 -1 2 1 0 -1
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF351E)
}
