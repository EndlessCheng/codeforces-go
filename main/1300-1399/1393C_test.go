package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1393/C
// https://codeforces.com/problemset/status/1393/problem/C
func TestCF1393C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7
1 7 1 6 4 4 6
8
1 1 4 6 4 6 4 7
3
3 3 3
6
2 5 2 3 1 4
outputCopy
3
2
0
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1393C)
}
