package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/846/C
// https://codeforces.com/problemset/status/846/problem/C
func TestCF846C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
-1 2 3
outputCopy
0 1 3
inputCopy
4
0 0 -1 0
outputCopy
0 0 0
inputCopy
1
10000
outputCopy
1 1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF846C)
}
