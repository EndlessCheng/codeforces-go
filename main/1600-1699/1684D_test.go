package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1684/D
// https://codeforces.com/problemset/status/1684/problem/D
func TestCF1684D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 4
8 7 1 4
4 1
5 10 11 5
7 5
8 2 5 15 11 2 8
6 3
1 2 3 4 5 6
1 1
7
outputCopy
0
21
9
6
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1684D)
}
