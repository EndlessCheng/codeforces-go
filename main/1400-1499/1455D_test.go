package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1455/D
// https://codeforces.com/problemset/status/1455/problem/D
func TestCF1455D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 1
2 3 5 4
5 6
1 1 3 4 4
1 10
2
2 10
11 9
2 10
12 11
5 18
81 324 218 413 324
outputCopy
3
0
0
-1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1455D)
}
