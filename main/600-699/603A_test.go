package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/603/A
// https://codeforces.com/problemset/status/603/problem/A
func TestCF603A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
10000011
outputCopy
5
inputCopy
2
01
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF603A)
}
