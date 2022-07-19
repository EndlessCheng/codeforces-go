package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/76/D
// https://codeforces.com/problemset/status/76/problem/D
func TestCF76D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
142
76
outputCopy
33 109`
	testutil.AssertEqualCase(t, rawText, 0, CF76D)
}
