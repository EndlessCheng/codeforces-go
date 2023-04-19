package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1630/problem/C
// https://codeforces.com/problemset/status/1630/problem/C
func TestCF1630C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 2 1 2 7 4 7
outputCopy
2
inputCopy
13
1 2 3 2 1 3 3 4 5 5 5 4 7
outputCopy
7
inputCopy
3
1 2 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1630C)
}
