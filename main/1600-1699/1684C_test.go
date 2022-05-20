package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1684/C
// https://codeforces.com/problemset/status/1684/problem/C
func TestCF1684C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 3
1 2 3
1 1 1
2 2
4 1
2 3
2 2
2 1
1 1
2 3
6 2 1
5 4 3
2 1
1
2
outputCopy
1 1
-1
1 2
1 3
1 1
inputCopy
1
1 7
1 2 4 3 3 3 5
outputCopy
3 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1684C)
}
