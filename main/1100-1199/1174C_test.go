package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1174/problem/C
// https://codeforces.com/problemset/status/1174/problem/C
func TestCF1174C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
1 2 1 
inputCopy
3
outputCopy
2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1174C)
}
