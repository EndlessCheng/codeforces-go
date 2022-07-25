package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1119/C
// https://codeforces.com/problemset/status/1119/problem/C
func TestCF1119C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
0 1 0
0 1 0
1 0 0
1 0 0
1 0 0
1 0 0
outputCopy
Yes
inputCopy
6 7
0 0 1 1 0 0 1
0 1 0 0 1 0 1
0 0 0 1 0 0 1
1 0 1 0 1 0 0
0 1 0 0 1 0 1
0 1 0 1 0 0 1
1 1 0 1 0 1 1
0 1 1 0 1 0 0
1 1 0 1 0 0 1
1 0 1 0 0 1 0
0 1 1 0 1 0 0
0 1 1 1 1 0 1
outputCopy
Yes
inputCopy
3 4
0 1 0 1
1 0 1 0
0 1 0 1
1 1 1 1
1 1 1 1
1 1 1 1
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1119C)
}
