package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/237/problem/C
// https://codeforces.com/problemset/status/237/problem/C
func TestCF237C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4 2
outputCopy
3
inputCopy
6 13 1
outputCopy
4
inputCopy
1 4 3
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF237C)
}
