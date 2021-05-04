package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/356/C
// https://codeforces.com/problemset/status/356/problem/C
func TestCF356C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 2 4 3
outputCopy
2
inputCopy
3
4 1 1
outputCopy
2
inputCopy
4
0 3 0 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF356C)
}
