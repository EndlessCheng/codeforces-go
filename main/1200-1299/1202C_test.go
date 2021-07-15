package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1202/problem/C
// https://codeforces.com/problemset/status/1202/problem/C
func TestCF1202C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
DSAWWAW
D
WA
outputCopy
8
2
4
inputCopy
1
WA
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, -1, CF1202C)
}
