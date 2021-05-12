package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/235/B
// https://codeforces.com/problemset/status/235/problem/B
func TestCF235B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0.5 0.5 0.5
outputCopy
2.750000000000000
inputCopy
4
0.7 0.2 0.1 0.9
outputCopy
2.489200000000000
inputCopy
5
1 1 1 1 1
outputCopy
25.000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF235B)
}
