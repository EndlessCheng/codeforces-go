package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/305/C
// https://codeforces.com/problemset/status/305/problem/C
func TestCF305C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 1 1 1
outputCopy
0
inputCopy
1
3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF305C)
}
