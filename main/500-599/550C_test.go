package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/550/C
// https://codeforces.com/problemset/status/550/problem/C
func TestCF550C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3454
outputCopy
YES
344
inputCopy
10
outputCopy
YES
0
inputCopy
111111
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF550C)
}
