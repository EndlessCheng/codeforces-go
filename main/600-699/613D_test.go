package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/613/D
// https://codeforces.com/problemset/status/613/problem/D
func TestCF613D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 3
2 3
4 3
4
2 1 2
3 2 3 4
3 1 2 4
4 1 2 3 4
outputCopy
1
-1
1
-1
inputCopy
7
1 2
2 3
3 4
1 5
5 6
5 7
1
4 2 4 6 7
outputCopy
2
inputCopy
4
1 2
2 3
1 4
1
3 1 3 4
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF613D)
}
