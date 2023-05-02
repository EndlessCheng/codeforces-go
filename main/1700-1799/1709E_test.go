package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1709/E
// https://codeforces.com/problemset/status/1709/problem/E
func TestCF1709E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 2 1 3 2 1
4 5
3 4
1 4
2 1
6 1
outputCopy
2
inputCopy
4
2 1 1 1
1 2
1 3
1 4
outputCopy
0
inputCopy
5
2 2 2 2 2
1 2
2 3
3 4
4 5
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1709E)
}
