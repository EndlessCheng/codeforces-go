package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/75/D
// https://codeforces.com/problemset/status/75/problem/D
func TestCF75D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
3 1 6 -2
2 3 3
2 -5 1
2 3 1 3
outputCopy
9
inputCopy
6 1
4 0 8 -3 -10
8 3 -2 -5 10 8 -9 -5 -4
1 0
1 -3
3 -8 5 6
2 9 6
1
outputCopy
8
inputCopy
3 6
3 -1 -1 -1
4 -2 -2 -2 -2
5 -3 -3 -3 -3 -3
1 2 3 1 2 3
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 1, CF75D)
}
