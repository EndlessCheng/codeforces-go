package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1624/problem/C
// https://codeforces.com/problemset/status/1624/problem/C
func TestCF1624C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
1 8 25 2
2
1 1
9
9 8 3 4 2 7 1 5 6
3
8 2 1
4
24 7 16 7
5
22 6 22 4 22
outputCopy
YES
NO
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1624C)
}
