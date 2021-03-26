package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/812/E
// https://codeforces.com/problemset/status/812/problem/E
func TestCF812E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 2 3
1 1
outputCopy
1
inputCopy
3
1 2 3
1 1
outputCopy
0
inputCopy
8
7 2 2 5 4 3 1 1
1 1 1 4 4 5 6
outputCopy
4
inputCopy
11
1 1 1 1 1 1 1 1 1 1 1
1 2 3 4 5 6 6 6 6 6
outputCopy
55`
	testutil.AssertEqualCase(t, rawText, -1, CF812E)
}
