package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1467/B
// https://codeforces.com/problemset/status/1467/problem/B
func TestCF1467B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 5 3
5
2 2 2 2 2
6
1 6 2 5 2 10
5
1 6 2 5 1
outputCopy
0
0
1
0
inputCopy
1
8
23 28 26 23 28 5 16 16
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1467B)
}
