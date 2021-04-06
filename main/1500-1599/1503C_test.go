package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1503/C
// https://codeforces.com/problemset/status/1503/problem/C
func TestCF1503C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 9
2 1
4 1
outputCopy
11
inputCopy
6
4 2
8 4
3 0
2 3
7 1
0 1
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, 0, CF1503C)
}
