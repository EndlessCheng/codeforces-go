package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1517/B
// https://codeforces.com/problemset/status/1517/problem/B
func TestCF1517B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 3
2 3 4
1 3 5
3 2
2 3
4 1
3 5
outputCopy
2 3 4
5 3 1
2 3
4 1
3 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1517B)
}
