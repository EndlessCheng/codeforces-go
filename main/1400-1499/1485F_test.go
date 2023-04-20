package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1485/F
// https://codeforces.com/problemset/status/1485/problem/F
func TestCF1485F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 -1 1
4
1 2 3 4
10
2 -1 1 -2 2 3 -5 0 2 -1
4
0 0 0 1
outputCopy
3
8
223
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1485F)
}
