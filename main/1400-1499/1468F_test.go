package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1468/F
// https://codeforces.com/problemset/status/1468/problem/F
func TestCF1468F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
0 0 0 1
1 0 2 0
3
0 0 1 1
1 1 0 0
1 0 2 0
6
0 0 0 1
1 0 1 2
2 0 2 3
3 0 3 -5
4 0 4 -5
5 0 5 -5
outputCopy
0
1
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1468F)
}
