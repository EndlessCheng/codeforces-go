package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1528/C
// https://codeforces.com/problemset/status/1528/problem/C
func TestCF1528C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
1 2 3
1 2 3
5
1 2 3 4
1 1 1 1
6
1 1 1 1 2
1 2 1 2 2
7
1 1 3 4 4 5
1 2 1 4 2 5
outputCopy
1
4
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1528C)
}
