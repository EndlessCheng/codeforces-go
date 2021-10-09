package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/7/D
// https://codeforces.com/problemset/status/7/problem/D
func TestCF7D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
a2A
outputCopy
1
inputCopy
abacaba
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF7D)
}
