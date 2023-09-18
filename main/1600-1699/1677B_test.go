package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1677/B
// https://codeforces.com/problemset/status/1677/problem/B
func TestCF1677B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2
1100
4 2
11001101
2 4
11001101
outputCopy
2 3 4 3
2 3 4 3 5 4 6 5
2 3 3 3 4 4 4 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1677B)
}
