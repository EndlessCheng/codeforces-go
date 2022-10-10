package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 单调队列上三分？

// https://codeforces.com/problemset/problem/988/F
// https://codeforces.com/problemset/status/988/problem/F
func TestCF988F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 2 4
3 7
8 10
0 10
3 4
8 1
1 2
outputCopy
14
inputCopy
10 1 1
0 9
0 5
outputCopy
45
inputCopy
10 1 1
0 9
1 5
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF988F)
}
