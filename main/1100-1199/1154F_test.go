package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1154/F
// https://codeforces.com/problemset/status/1154/problem/F
func TestCF1154F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 4 5
2 5 4 2 6 3 1
2 1
6 5
2 1
3 1
outputCopy
7
inputCopy
9 4 8
6 8 5 1 8 1 1 2 1
9 2
8 4
5 3
9 7
outputCopy
17
inputCopy
5 1 4
2 5 7 4 6
5 4
outputCopy
17`
	testutil.AssertEqualCase(t, rawText, 0, CF1154F)
}
