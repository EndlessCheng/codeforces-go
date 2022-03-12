package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1650/problem/B
// https://codeforces.com/problemset/status/1650/problem/B
func TestCF1650B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 4 3
5 8 4
6 10 6
1 1000000000 1000000000
10 12 8
outputCopy
2
4
5
999999999
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1650B)
}
