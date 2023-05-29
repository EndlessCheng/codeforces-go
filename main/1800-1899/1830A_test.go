package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1830/problem/A
// https://codeforces.com/problemset/status/1830/problem/A
func TestCF1830A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
6
4 5
1 3
1 2
3 4
1 6
7
5 6
2 4
2 7
1 3
1 2
4 5
outputCopy
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1830A)
}
