package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1479/B1
// https://codeforces.com/problemset/status/1479/problem/B1
func TestCF1479B1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 1 2 2 3 3 3
outputCopy
6
inputCopy
7
1 2 3 4 5 6 7
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1479B1)
}
