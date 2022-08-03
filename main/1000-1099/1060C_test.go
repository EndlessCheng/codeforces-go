package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1060/C
// https://codeforces.com/problemset/status/1060/problem/C
func TestCF1060C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2 3
1 2 3
9
outputCopy
4
inputCopy
5 1
5 4 2 4 5
2
5
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1060C)
}
