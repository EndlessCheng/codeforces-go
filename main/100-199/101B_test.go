package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/101/B
// https://codeforces.com/problemset/status/101/problem/B
func TestCF101B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
0 1
1 2
outputCopy
1
inputCopy
3 2
0 1
1 2
outputCopy
0
inputCopy
5 5
0 1
0 2
0 3
0 4
0 5
outputCopy
16
inputCopy
7 7
0 1
1 3
2 3
4 6
5 7
4 5
5 7
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF101B)
}
