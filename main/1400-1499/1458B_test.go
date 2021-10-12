package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1458/B
// https://codeforces.com/problemset/status/1458/problem/B
func TestCF1458B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6 5
6 5
10 2
outputCopy
7.0000000000 11.0000000000 12.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1458B)
}
