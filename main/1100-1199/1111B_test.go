package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1111/B
// https://codeforces.com/problemset/status/1111/problem/B
func TestCF1111B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4 6
4 7
outputCopy
11.00000000000000000000
inputCopy
4 2 6
1 3 2 3
outputCopy
5.00000000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1111B)
}
