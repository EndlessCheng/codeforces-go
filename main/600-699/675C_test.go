package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/675/C
// https://codeforces.com/problemset/status/675/problem/C
func TestCF675C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 0 -5
outputCopy
1
inputCopy
4
-1 0 1 0
outputCopy
2
inputCopy
4
1 2 3 -6
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF675C)
}
