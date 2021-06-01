package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/607/A
// https://codeforces.com/problemset/status/607/problem/A
func TestCF607A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 9
3 1
6 1
7 4
outputCopy
1
inputCopy
7
1 1
2 1
3 1
4 1
5 1
6 1
7 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF607A)
}
