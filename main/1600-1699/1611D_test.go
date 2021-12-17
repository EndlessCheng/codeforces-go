package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1611/D
// https://codeforces.com/problemset/status/1611/problem/D
func TestCF1611D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
3 1 3 3 1
3 1 2 5 4
3
1 1 2
3 1 2
7
1 1 2 3 4 5 6
1 2 3 4 5 6 7
6
4 4 4 4 1 1
4 2 1 5 6 3
outputCopy
1 10 0 102 100
-1
0 3 100 1 1 2 4
6 5 10 0 2 3
inputCopy
1
9
2 2 2 2 2 2 2 2 2
1 2 6 7 3 9 5 8 4
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, -1, CF1611D)
}
