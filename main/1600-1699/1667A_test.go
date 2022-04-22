package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1667/A
// https://codeforces.com/problemset/status/1667/problem/A
func TestCF1667A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
outputCopy
4
inputCopy
7
1 2 1 2 1 2 1
outputCopy
10
inputCopy
8
1 8 2 7 3 6 4 5
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1667A)
}
