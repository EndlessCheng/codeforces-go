package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1428/E
// https://codeforces.com/problemset/status/1428/problem/E
func TestCF1428E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 6
5 3 1
outputCopy
15
inputCopy
1 4
19
outputCopy
91`
	testutil.AssertEqualCase(t, rawText, 0, CF1428E)
}
