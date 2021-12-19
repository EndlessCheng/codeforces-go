package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/682/A
// https://codeforces.com/problemset/status/682/problem/A
func TestCF682A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 12
outputCopy
14
inputCopy
11 14
outputCopy
31
inputCopy
1 5
outputCopy
1
inputCopy
3 8
outputCopy
5
inputCopy
5 7
outputCopy
7
inputCopy
21 21
outputCopy
88`
	testutil.AssertEqualCase(t, rawText, 0, CF682A)
}
