package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1519/problem/D
// https://codeforces.com/problemset/status/1519/problem/D
func TestCF1519D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 3 2 1 3
1 3 2 4 2
outputCopy
29
inputCopy
2
13 37
2 4
outputCopy
174
inputCopy
6
1 8 7 6 3 6
5 9 6 8 8 6
outputCopy
235`
	testutil.AssertEqualCase(t, rawText, 0, CF1519D)
}
