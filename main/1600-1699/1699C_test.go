package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1699/C
// https://codeforces.com/problemset/status/1699/problem/C
func TestCF1699C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
4 0 3 2 1
1
0
4
0 1 2 3
6
1 2 4 0 5 3
8
1 3 7 2 5 0 6 4
outputCopy
2
1
1
4
72`
	testutil.AssertEqualCase(t, rawText, 0, CF1699C)
}
