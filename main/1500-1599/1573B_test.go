package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1573/B
// https://codeforces.com/problemset/status/1573/problem/B
func TestCF1573B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
3 1
4 2
3
5 3 1
2 4 6
5
7 5 9 1 3
2 4 6 10 8
outputCopy
0
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1573B)
}
