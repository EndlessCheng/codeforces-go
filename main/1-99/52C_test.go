package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/52/C
// https://codeforces.com/problemset/status/52/problem/C
func TestCF52C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 3 4
4
3 0
3 0 -1
0 1
2 1
outputCopy
1
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF52C)
}
