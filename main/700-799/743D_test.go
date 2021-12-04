package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/743/D
// https://codeforces.com/problemset/status/743/problem/D
func TestCF743D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
0 5 -1 4 3 2 6 5
1 2
2 4
2 5
1 3
3 6
6 7
6 8
outputCopy
25
inputCopy
4
1 -5 1 1
1 2
1 4
2 3
outputCopy
2
inputCopy
1
-1
outputCopy
Impossible`
	testutil.AssertEqualCase(t, rawText, 2, CF743D)
}
