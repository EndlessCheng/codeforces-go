package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1765/D
// https://codeforces.com/problemset/status/1765/problem/D
func TestCF1765D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
1 2 3 4 5
outputCopy
16
inputCopy
5 5
1 2 3 4 5
outputCopy
17
inputCopy
4 3
1 3 2 3
outputCopy
12
inputCopy
10 10
2 3 2 6 8 6 6 8 8 7
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, -1, CF1765D)
}
