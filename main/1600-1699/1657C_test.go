package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1657/problem/C
// https://codeforces.com/problemset/status/1657/problem/C
func TestCF1657C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2
()
3
())
4
((((
5
)((()
6
)((()(
outputCopy
1 0
1 1
2 0
1 0
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1657C)
}
