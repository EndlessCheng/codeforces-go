package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1036/problem/D
// https://codeforces.com/problemset/status/1036/problem/D
func TestCF1036D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
11 2 3 5 7
4
11 7 3 7
outputCopy
3
inputCopy
2
1 2
1
100
outputCopy
-1
inputCopy
3
1 2 3
3
1 2 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1036D)
}
