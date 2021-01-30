package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/E2
// https://codeforces.com/problemset/status/1399/problem/E2
func TestCF1399E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 18
2 1 9 2
3 2 4 1
4 1 1 2
3 20
2 1 8 1
3 1 7 2
5 50
1 3 100 1
1 5 10 2
2 3 123 2
5 4 55 1
2 100
1 2 409 2
outputCopy
0
0
11
6`
	testutil.AssertEqualCase(t, rawText, -1, CF1399E2)
}
