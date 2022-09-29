package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/558/C
// https://codeforces.com/problemset/status/558/problem/C
func TestCF558C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 8 2
outputCopy
2
inputCopy
3
3 5 6
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF558C)
}
