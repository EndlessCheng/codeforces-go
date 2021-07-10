package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/21/C
// https://codeforces.com/problemset/status/21/problem/C
func TestCF21C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 3 3
outputCopy
1
inputCopy
5
1 2 3 4 5
outputCopy
0
inputCopy
4
0 0 0 0
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF21C)
}
