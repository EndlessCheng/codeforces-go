package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1365/F
// https://codeforces.com/problemset/status/1365/problem/F
func TestCF1365F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2
1 2
2 1
3
1 2 3
1 2 3
3
1 2 4
1 3 4
4
1 2 3 2
3 1 2 2
3
1 2 3
1 3 2
outputCopy
Yes
Yes
No
Yes
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1365F)
}
