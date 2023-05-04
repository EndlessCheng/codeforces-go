package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1774/E
// https://codeforces.com/problemset/status/1774/problem/E
func TestCF1774E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2
1 3
2 4
1 3
1 4
outputCopy
6
inputCopy
4 2
1 2
2 3
3 4
4 1 2 3 4
1 1
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1774E)
}
