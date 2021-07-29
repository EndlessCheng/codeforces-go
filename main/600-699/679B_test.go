package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/679/B
// https://codeforces.com/problemset/status/679/problem/B
func TestCF679B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
48
outputCopy
9 42
inputCopy
6
outputCopy
6 6`
	testutil.AssertEqualCase(t, rawText, 0, CF679B)
}
