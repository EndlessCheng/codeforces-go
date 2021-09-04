package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/808/E
// https://codeforces.com/problemset/status/808/problem/E
func TestCF808E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
2 1
outputCopy
0
inputCopy
2 2
1 3
2 2
outputCopy
3
inputCopy
4 3
3 10
2 7
2 8
1 1
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF808E)
}
