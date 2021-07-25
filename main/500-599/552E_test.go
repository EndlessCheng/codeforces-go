package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/552/E
// https://codeforces.com/problemset/status/552/problem/E
func TestCF552E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3+5*7+8*4
outputCopy
303
inputCopy
2+3*5
outputCopy
25
inputCopy
3*4*5
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, 1, CF552E)
}
