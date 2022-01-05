package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1621/B
// https://codeforces.com/problemset/status/1621/problem/B
func TestCF1621B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
2 4 20
7 8 22
2
5 11 42
5 11 42
6
1 4 4
5 8 9
7 8 7
2 10 252
1 11 271
1 10 1
outputCopy
20
42
42
42
4
13
11
256
271
271`
	testutil.AssertEqualCase(t, rawText, 0, CF1621B)
}
