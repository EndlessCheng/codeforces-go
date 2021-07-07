package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/792/C
// https://codeforces.com/problemset/status/792/problem/C
func TestCF792C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1033
outputCopy
33
inputCopy
10
outputCopy
0
inputCopy
11
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF792C)
}
