package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1689/E
// https://codeforces.com/problemset/status/1689/problem/E
func TestCF1689E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 2 3 4 5
2
0 2
2
3 12
4
3 0 0 0
outputCopy
0
1 2 3 4 5
2
2 2
1
3 11
3
3 1 1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1689E)
}
