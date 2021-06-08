package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/219/C
// https://codeforces.com/problemset/status/219/problem/C
func TestCF219C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
ABBACC
outputCopy
2
ABCACA
inputCopy
3 2
BBB
outputCopy
1
BAB`
	testutil.AssertEqualCase(t, rawText, 0, CF219C)
}
