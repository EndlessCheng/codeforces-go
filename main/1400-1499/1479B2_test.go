package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1479/B2
// https://codeforces.com/problemset/status/1479/problem/B2
func TestCF1479B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2 3 1 2 2
outputCopy
4
inputCopy
7
1 2 1 2 1 2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1479B2)
}
