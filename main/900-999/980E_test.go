package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/980/E
// https://codeforces.com/problemset/status/980/problem/E
func TestCF980E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
2 1
2 6
4 2
5 6
2 3
outputCopy
1 3 4
inputCopy
8 4
2 6
2 7
7 8
1 2
3 1
2 4
7 5
outputCopy
1 3 4 5`
	testutil.AssertEqualCase(t, rawText, 0, CF980E)
}
