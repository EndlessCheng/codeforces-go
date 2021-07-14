package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/652/problem/D
// https://codeforces.com/problemset/status/652/problem/D
func TestCF652D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 8
2 3
4 7
5 6
outputCopy
3
0
1
0
inputCopy
3
3 4
1 5
2 6
outputCopy
0
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF652D)
}
