package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1114/D
// https://codeforces.com/problemset/status/1114/problem/D
func TestCF1114D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 2 2 1
outputCopy
2
inputCopy
8
4 5 2 2 1 3 5 5
outputCopy
4
inputCopy
1
4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1114D)
}
