package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1697/problem/C
// https://codeforces.com/problemset/status/1697/problem/C
func TestCF1697C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
cab
cab
1
a
b
6
abbabc
bbaacb
10
bcaabababc
cbbababaac
2
ba
ab
outputCopy
YES
NO
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1697C)
}
