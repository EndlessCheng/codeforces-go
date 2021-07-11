package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/258/C
// https://codeforces.com/problemset/status/258/problem/C
func TestCF258C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 4 3 2
outputCopy
15
inputCopy
2
6 3
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, 1, CF258C)
}
