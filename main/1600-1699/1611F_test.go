package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1611/F
// https://codeforces.com/problemset/status/1611/problem/F
func TestCF1611F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 10
-16 2 -6 8
3 1000
-100000 -100000 -100000
6 0
2 6 -164 1 -1 -6543
outputCopy
2 4
-1
1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1611F)
}
