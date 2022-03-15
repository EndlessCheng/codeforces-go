package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1618/problem/D
// https://codeforces.com/problemset/status/1618/problem/D
func TestCF1618D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
7 3
1 1 1 2 1 3 1
5 1
5 5 5 5 5
4 2
1 3 3 7
2 0
4 2
9 2
1 10 10 1 10 2 7 10 3
outputCopy
2
16
0
6
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1618D)
}
