package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/954/problem/C
// https://codeforces.com/problemset/status/954/problem/C
func TestCF954C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
1 2 3 6 9 8 5 2
outputCopy
YES
3 3
inputCopy
6
1 2 1 2 5 3
outputCopy
NO
inputCopy
2
1 10
outputCopy
YES
4 9`
	testutil.AssertEqualCase(t, rawText, 0, CF954C)
}
