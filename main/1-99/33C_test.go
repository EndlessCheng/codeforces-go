package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/33/C
// https://codeforces.com/problemset/status/33/problem/C
func TestCF33C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
-1 -2 -3
outputCopy
6
inputCopy
5
-4 2 0 5 0
outputCopy
11
inputCopy
5
-1 10 -5 10 -2
outputCopy
18`
	testutil.AssertEqualCase(t, rawText, 0, CF33C)
}
