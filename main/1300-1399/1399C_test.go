package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/C
// https://codeforces.com/problemset/status/1399/problem/C
func TestCF1399C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 2 3 4 5
8
6 6 6 6 6 6 8 8
8
1 2 2 1 2 1 1 2
3
1 3 3
6
1 1 3 4 2 2
outputCopy
2
3
4
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1399C)
}
