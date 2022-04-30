package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1622/problem/C
// https://codeforces.com/problemset/status/1622/problem/C
func TestCF1622C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 10
20
2 69
6 9
7 8
1 2 1 3 1 2 1
10 1
1 2 3 1 2 6 1 6 8 10
outputCopy
10
0
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1622C)
}
