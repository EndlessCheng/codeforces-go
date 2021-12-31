package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1366/E
// https://codeforces.com/problemset/status/1366/problem/E
func TestCF1366E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
12 10 20 20 25 30
10 20 30
outputCopy
2
inputCopy
4 2
1 3 3 7
3 7
outputCopy
0
inputCopy
8 2
1 2 2 2 2 2 2 2
1 2
outputCopy
7
inputCopy
6 3
12 10 20 20 15 30
10 20 30
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1366E)
}
