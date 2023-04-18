package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1344/A
// https://codeforces.com/problemset/status/1344/problem/A
func TestCF1344A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1
14
2
1 -1
4
5 5 5 1
3
3 2 1
2
0 1
5
-239 -2 -100 -3 -11
outputCopy
YES
YES
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1344A)
}
