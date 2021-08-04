package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/520/E
// https://codeforces.com/problemset/status/520/problem/E
func TestCF520E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
108
outputCopy
27
inputCopy
3 2
108
outputCopy
9
inputCopy
1 0
5
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF520E)
}
