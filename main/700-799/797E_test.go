package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/797/E
// https://codeforces.com/problemset/status/797/problem/E
func TestCF797E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1 1
3
1 1
2 1
3 1
outputCopy
2
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF797E)
}
