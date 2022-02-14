package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1238/problem/F
// https://codeforces.com/problemset/status/1238/problem/F
func TestCF1238F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
10
1 2
1 3
1 4
2 5
2 6
3 7
3 8
4 9
4 10
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1238F)
}
