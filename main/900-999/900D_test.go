package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/900/D
// https://codeforces.com/problemset/status/900/problem/D
func TestCF900D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 9
outputCopy
3
inputCopy
5 8
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF900D)
}
