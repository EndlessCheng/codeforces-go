package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1558/C
// https://codeforces.com/problemset/status/1558/problem/C
func TestCF1558C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 2 3
5
3 4 5 2 1
3
2 1 3
outputCopy
4
3 3 3 3
2
3 5
-1
inputCopy
1
5
3 4 1 2 5
outputCopy
10
5 3 5 3 5 3 1 3 3 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1558C)
}
