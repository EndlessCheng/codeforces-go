package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1263/problem/F
// https://codeforces.com/problemset/status/1263/problem/F
func TestCF1263F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
4 1 1 4 2
6 5 3
4
1 1 1
3 4 2
outputCopy
5
inputCopy
4
6
4 4 1 1 1
3 2 6 5
6
6 6 1 1 1
5 4 3 2
outputCopy
6
inputCopy
5
14
1 1 11 2 14 14 13 7 12 2 5 6 1
9 8 3 10 4
16
1 1 9 9 2 5 10 1 14 3 7 11 6 12 2
8 16 13 4 15
outputCopy
17`
	testutil.AssertEqualCase(t, rawText, 0, CF1263F)
}
