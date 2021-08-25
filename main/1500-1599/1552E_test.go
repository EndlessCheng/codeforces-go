package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1552/E
// https://codeforces.com/problemset/status/1552/problem/E
func TestCF1552E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
2 4 3 1 1 4 2 3 2 1 3 4
outputCopy
4 5
1 7
8 11
6 12
inputCopy
1 2
1 1
outputCopy
1 2
inputCopy
3 3
3 1 2 3 2 1 2 1 3
outputCopy
6 8
3 7
1 4
inputCopy
2 3
2 1 1 1 2 2
outputCopy
2 3
5 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1552E)
}
