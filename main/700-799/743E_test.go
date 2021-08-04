package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/743/E
// https://codeforces.com/problemset/status/743/problem/E
func TestCF743E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1 1
outputCopy
1
inputCopy
8
8 7 6 5 4 3 2 1
outputCopy
8
inputCopy
24
1 8 1 2 8 2 3 8 3 4 8 4 5 8 5 6 8 6 7 8 7 8 8 8
outputCopy
17`
	testutil.AssertEqualCase(t, rawText, 0, CF743E)
}
