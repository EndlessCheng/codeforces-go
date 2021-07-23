package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/372/C
// https://codeforces.com/problemset/status/372/problem/C
func TestCF372C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
50 3 1
49 1 1
26 1 4
6 1 10
outputCopy
-31
inputCopy
10 2 1
1 1000 4
9 1000 4
outputCopy
1992`
	testutil.AssertEqualCase(t, rawText, 0, CF372C)
}
