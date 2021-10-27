package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1000/B
// https://codeforces.com/problemset/status/1000/problem/B
func TestCF1000B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 10
4 6 7
outputCopy
8
inputCopy
2 12
1 10
outputCopy
9
inputCopy
2 7
3 4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1000B)
}
