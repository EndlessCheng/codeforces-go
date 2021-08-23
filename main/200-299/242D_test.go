package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/242/D
// https://codeforces.com/problemset/status/242/problem/D
func TestCF242D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
2 3
4 1
1 5
5 3
2 1
1 1 2 0 2
outputCopy
2
1 2
inputCopy
4 2
1 2
3 4
0 0 0 0
outputCopy
3
1 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF242D)
}
