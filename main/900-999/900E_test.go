package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/900/E
// https://codeforces.com/problemset/status/900/problem/E
func TestCF900E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
bb?a?
1
outputCopy
2
inputCopy
9
ab??ab???
3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF900E)
}
