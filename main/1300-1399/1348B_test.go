package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1348/B
// https://codeforces.com/problemset/status/1348/problem/B
func TestCF1348B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 2
1 2 2 1
4 3
1 2 2 1
3 2
1 2 3
4 4
4 3 4 2
outputCopy
5
1 2 1 2 1
4
1 2 2 1
-1
7
4 3 2 1 4 3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1348B)
}
