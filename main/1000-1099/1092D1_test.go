package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1092/D1
// https://codeforces.com/problemset/status/1092/problem/D1
func TestCF1092D1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1 1 2 5
outputCopy
YES
inputCopy
3
4 5 3
outputCopy
YES
inputCopy
2
10 10
outputCopy
YES
inputCopy
3
1 2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1092D1)
}
