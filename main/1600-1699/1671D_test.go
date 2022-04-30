package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1671/D
// https://codeforces.com/problemset/status/1671/problem/D
func TestCF1671D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 5
10
3 8
7 2 10
10 2
6 1 5 7 3 3 9 10 10 1
4 10
1 3 1 2
outputCopy
9
15
31
13`
	testutil.AssertEqualCase(t, rawText, 0, CF1671D)
}
