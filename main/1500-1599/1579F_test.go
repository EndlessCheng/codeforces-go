package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1579/F
// https://codeforces.com/problemset/status/1579/problem/F
func TestCF1579F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1
0 1
3 2
0 1 0
5 2
1 1 0 1 0
4 2
0 1 0 1
1 1
0
outputCopy
1
1
3
-1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1579F)
}
