package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1093/G
// https://codeforces.com/problemset/status/1093/problem/G
func TestCF1093G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
1 2
2 3
3 4
4 5
5 6
7
2 1 5
2 1 3
2 3 5
1 5 -1 -2
2 1 5
1 4 -1 -2
2 1 5
outputCopy
8
4
4
12
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1093G)
}
