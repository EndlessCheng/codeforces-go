package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1178/problem/D
// https://codeforces.com/problemset/status/1178/problem/D
func TestCF1178D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
5
1 2
1 3
2 3
2 4
3 4
inputCopy
8
outputCopy
13
1 2
1 3
2 3
1 4
2 4
1 5
2 5
1 6
2 6
1 7
1 8
5 8
7 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1178D)
}
