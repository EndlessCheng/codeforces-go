package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1380/C
// https://codeforces.com/problemset/status/1380/problem/C
func TestCF1380C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 10
7 11 2 9 5
4 8
2 4 2 3
4 11
1 3 3 7
outputCopy
2
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1380C)
}
