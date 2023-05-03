package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1788/F
// https://codeforces.com/problemset/status/1788/problem/F
func TestCF1788F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2
2 3
3 4
1 4 3
2 4 2
1 3 1
2 3 1
outputCopy
No
inputCopy
6 2
1 2
2 3
3 4
2 5
5 6
1 4 2
2 6 7
outputCopy
Yes
4 2 4 1 6
inputCopy
6 2
1 2
2 3
3 4
2 5
5 6
1 4 3
1 6 5
outputCopy
Yes
6 1 4 3 0`
	testutil.AssertEqualCase(t, rawText, 0, CF1788F)
}
