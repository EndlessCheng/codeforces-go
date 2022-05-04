package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1671/C
// https://codeforces.com/problemset/status/1671/problem/C
func TestCF1671C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 7
2 1 2
5 9
10 20 30 40 50
1 1
1
2 1000
1 1
outputCopy
11
0
1
1500`
	testutil.AssertEqualCase(t, rawText, 0, CF1671C)
}
