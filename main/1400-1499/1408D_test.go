package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1408/D
// https://codeforces.com/problemset/status/1408/problem/D
func TestCF1408D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
0 0
2 3
outputCopy
3
inputCopy
2 3
1 6
6 1
10 1
1 10
7 7
outputCopy
4
inputCopy
1 2
0 0
0 0
0 0
outputCopy
1
inputCopy
7 3
0 8
3 8
2 7
0 10
5 5
7 0
3 5
6 6
3 11
11 5
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1408D)
}
