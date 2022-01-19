package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/544/B
// https://codeforces.com/problemset/status/544/problem/B
func TestCF544B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
outputCopy
YES
SSSSS
LLLLL
SSSSS
LLLLL
SSSSS
inputCopy
5 25
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF544B)
}
