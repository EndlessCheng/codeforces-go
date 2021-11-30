package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1609/C
// https://codeforces.com/problemset/status/1609/problem/C
func TestCF1609C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
7 3
10 2 1 3 1 19 3
3 2
1 13 1
9 3
2 4 2 1 1 1 1 4 2
3 1
1 1 1
4 1
1 2 1 1
2 2
1 2
outputCopy
2
0
4
0
5
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1609C)
}
