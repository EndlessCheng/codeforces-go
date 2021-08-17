package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/788/C
// https://codeforces.com/problemset/status/788/problem/C
func TestCF788C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
400 4
100 300 450 500
outputCopy
2
inputCopy
50 2
100 25
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF788C)
}
