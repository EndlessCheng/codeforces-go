package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/743/C
// https://codeforces.com/problemset/status/743/problem/C
func TestCF743C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
2 7 42
inputCopy
7
outputCopy
7 8 56`
	testutil.AssertEqualCase(t, rawText, 0, CF743C)
}
