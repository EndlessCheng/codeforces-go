package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1392/D
// https://codeforces.com/problemset/status/1392/problem/D
func TestCF1392D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
RLRL
6
LRRRRL
8
RLLRRRLL
12
LLLLRRLRRRLL
5
RRRRR
outputCopy
0
1
1
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1392D)
}
