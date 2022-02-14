package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1093/problem/D
// https://codeforces.com/problemset/status/1093/problem/D
func TestCF1093D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 1
1 2
4 6
1 2
1 3
1 4
2 3
2 4
3 4
outputCopy
4
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1093D)
}
