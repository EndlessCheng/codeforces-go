package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/status/883/problem/I
func TestCF883I(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
50 110 130 40 120
outputCopy
20
inputCopy
4 1
2 3 4 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF883I)
}
