package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1384/B2
// https://codeforces.com/problemset/status/1384/problem/B2
func TestCF1384B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
2 1 1
1 0
5 2 3
1 2 3 2 2
4 3 4
0 2 4 3
2 3 5
3 0
7 2 3
3 0 2 1 3 0 1
7 1 4
4 4 3 0 2 4 2
5 2 3
1 2 3 2 2
outputCopy
Yes
No
Yes
Yes
Yes
No
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1384B2)
}
