package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1200/problem/F
// https://codeforces.com/problemset/status/1200/problem/F
func TestCF1200F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 0 0 0
2
2 3
1
2
3
2 4 1
4
3 1 2 1
6
1 0
2 0
3 -1
4 -2
1 1
1 5
outputCopy
1
1
2
1
3
2
inputCopy
4
4 -5 -3 -1
2
2 3
1
2
3
2 4 1
4
3 1 2 1
6
1 0
2 0
3 -1
4 -2
1 1
1 5
outputCopy
1
1
1
3
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1200F)
}
