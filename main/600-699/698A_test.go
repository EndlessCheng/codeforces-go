package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/698/A
// https://codeforces.com/problemset/status/698/problem/A
func TestCF698A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 3 2 0
outputCopy
2
inputCopy
7
1 3 3 2 1 2 3
outputCopy
0
inputCopy
2
2 2
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF698A)
}
