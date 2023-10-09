package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1882/problem/B
// https://codeforces.com/problemset/status/1882/problem/B
func TestCF1882B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
3 1 2 3
2 4 5
2 3 4
4
4 1 2 3 4
3 2 5 6
3 3 5 6
3 4 5 6
5
1 1
3 3 6 10
1 9
2 1 3
3 5 8 9
1
2 4 28
outputCopy
4
5
6
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1882B)
}
