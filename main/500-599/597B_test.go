package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/597/B
// https://codeforces.com/problemset/status/597/problem/B
func TestCF597B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
7 11
4 7
outputCopy
1
inputCopy
5
1 2
2 3
3 4
4 5
5 6
outputCopy
3
inputCopy
6
4 8
1 5
4 7
2 5
1 3
6 8
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF597B)
}
