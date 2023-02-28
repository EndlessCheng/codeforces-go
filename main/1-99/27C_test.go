package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/27/C
// https://codeforces.com/problemset/status/27/problem/C
func TestCF27C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
67 499 600 42 23
outputCopy
3
1 3 5
inputCopy
3
1 2 3
outputCopy
0
inputCopy
3
2 3 1
outputCopy
3
1 2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF27C)
}
