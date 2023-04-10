package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1201/C
// https://codeforces.com/problemset/status/1201/problem/C
func TestCF1201C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 3 5
outputCopy
5
inputCopy
5 5
1 2 1 1 1
outputCopy
3
inputCopy
7 7
4 1 2 4 3 4 4
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1201C)
}
