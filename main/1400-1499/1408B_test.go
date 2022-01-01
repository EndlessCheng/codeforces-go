package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1408/B
// https://codeforces.com/problemset/status/1408/problem/B
func TestCF1408B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 1
0 0 0 1
3 1
3 3 3
11 3
0 1 2 2 3 3 3 4 4 4 4
5 3
1 2 3 4 5
9 4
2 2 3 5 7 11 13 13 17
10 7
0 1 1 2 3 3 4 5 5 6
outputCopy
-1
1
2
2
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1408B)
}
