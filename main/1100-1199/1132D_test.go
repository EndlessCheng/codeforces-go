package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1132/problem/D
// https://codeforces.com/problemset/status/1132/problem/D
func TestCF1132D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4
3 2
4 2
outputCopy
5
inputCopy
1 5
4
2
outputCopy
1
inputCopy
1 6
4
2
outputCopy
2
inputCopy
2 2
2 10
3 15
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1132D)
}
