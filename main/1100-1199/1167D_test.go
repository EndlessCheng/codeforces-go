package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1167/D
// https://codeforces.com/problemset/status/1167/problem/D
func TestCF1167D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
()
outputCopy
11
inputCopy
4
(())
outputCopy
0101
inputCopy
10
((()())())
outputCopy
0110001111`
	testutil.AssertEqualCase(t, rawText, 0, CF1167D)
}
