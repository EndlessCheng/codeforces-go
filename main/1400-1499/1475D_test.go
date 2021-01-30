package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1475/problem/D
// https://codeforces.com/problemset/status/1475/problem/D
func TestCF1475D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 7
5 3 2 1 4
2 1 1 2 1
1 3
2
1
5 10
2 3 2 3 2
1 2 1 2 1
4 10
5 1 3 4
1 2 1 2
4 5
3 2 1 2
2 1 2 1
outputCopy
2
-1
6
4
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1475D)
}
