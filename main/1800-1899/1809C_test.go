package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1809/C
// https://codeforces.com/problemset/status/1809/problem/C
func TestCF1809C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 2
2 0
2 2
4 6
outputCopy
1 -3 1
-13 -42
-13 42
-3 -4 10 -2`
	testutil.AssertEqualCase(t, rawText, 0, CF1809C)
}
