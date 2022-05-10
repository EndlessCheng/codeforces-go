package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/865/D
// https://codeforces.com/problemset/status/865/problem/D
func TestCF865D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
10 5 4 7 9 12 6 2 10
outputCopy
20
inputCopy
20
3 1 4 1 5 9 2 6 5 3 5 8 9 7 9 3 2 3 8 4
outputCopy
41`
	testutil.AssertEqualCase(t, rawText, 0, CF865D)
}
