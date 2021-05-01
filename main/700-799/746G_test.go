package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/746/G
// https://codeforces.com/problemset/status/746/problem/G
func TestCF746G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3 3
2 3 1
outputCopy
7
1 3
2 1
2 6
2 4
7 4
3 5
inputCopy
14 5 6
4 4 2 2 1
outputCopy
14
3 1
1 4
11 6
1 2
10 13
6 10
10 12
14 12
8 4
5 1
3 7
2 6
5 9
inputCopy
3 1 1
2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 2, CF746G)
}
