package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1496/problem/B
// https://codeforces.com/problemset/status/1496/problem/B
func TestCF1496B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 1
0 1 3 4
3 1
0 1 4
3 0
0 1 4
3 2
0 1 2
3 2
1 2 3
outputCopy
4
4
3
5
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1496B)
}
