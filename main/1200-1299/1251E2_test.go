package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1251/E2
// https://codeforces.com/problemset/status/1251/problem/E2
func TestCF1251E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 5
2 10
2 8
7
0 1
3 1
1 1
6 1
1 1
4 1
4 1
6
2 6
2 3
2 8
2 7
4 4
5 5
outputCopy
8
0
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1251E2)
}
