package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/414/problem/B
// https://codeforces.com/problemset/status/414/problem/B
func TestCF414B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
outputCopy
5
inputCopy
6 4
outputCopy
39
inputCopy
2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF414B)
}
