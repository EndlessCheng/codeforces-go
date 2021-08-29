package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1560/B
// https://codeforces.com/problemset/status/1560/problem/B
func TestCF1560B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
6 2 4
2 3 1
2 4 10
5 3 4
1 3 2
2 5 4
4 3 2
outputCopy
8
-1
-1
-1
4
1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1560B)
}
