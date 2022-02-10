package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/8/B
// https://codeforces.com/problemset/status/8/problem/B
func TestCF8B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
LLUUUR
outputCopy
OK
inputCopy
RRUULLDD
outputCopy
BUG`
	testutil.AssertEqualCase(t, rawText, 0, CF8B)
}
