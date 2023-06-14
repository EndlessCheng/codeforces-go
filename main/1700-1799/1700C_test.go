package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1700/C
// https://codeforces.com/problemset/status/1700/problem/C
func TestCF1700C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
-2 -2 -2
3
10 4 7
4
4 -4 4 -4
5
1 -2 3 -4 5
outputCopy
2
13
36
33`
	testutil.AssertEqualCase(t, rawText, 0, CF1700C)
}
