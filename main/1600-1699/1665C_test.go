package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1665/C
// https://codeforces.com/problemset/status/1665/problem/C
func TestCF1665C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
7
1 1 1 2 2 4
5
5 5 1 4
2
1
3
3 1
6
1 1 1 1 1
outputCopy
4
4
2
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1665C)
}
