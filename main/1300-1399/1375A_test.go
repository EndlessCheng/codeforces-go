package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1375/A
// https://codeforces.com/problemset/status/1375/problem/A
func TestCF1375A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
-2 4 3
5
1 1 1 1 1
5
-2 4 7 -6 4
9
9 7 -4 -2 1 -3 9 -4 -5
9
-4 1 9 4 8 9 5 1 -9
outputCopy
-2 -4 3
1 1 1 1 1
-2 -4 7 -6 4
-9 -7 -4 2 1 -3 -9 -4 -5
4 -1 -9 -4 -8 -9 -5 -1 9`
	testutil.AssertEqualCase(t, rawText, 0, CF1375A)
}
