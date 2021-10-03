package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1540/A
// https://codeforces.com/problemset/status/1540/problem/A
func TestCF1540A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
0 2 3
2
0 1000000000
1
0
outputCopy
-3
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1540A)
}
