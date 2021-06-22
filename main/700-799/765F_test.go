package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/765/problem/F
// https://codeforces.com/problemset/status/765/problem/F
func TestCF765F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
3 1 4 1 5 9 2 6
4
1 8
1 3
4 8
5 7
outputCopy
0
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF765F)
}
