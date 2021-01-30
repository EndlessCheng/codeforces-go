package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/E1
// https://codeforces.com/problemset/status/1399/problem/E1
func TestCF1399E1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 20
2 1 8
3 1 7
5 50
1 3 100
1 5 10
2 3 123
5 4 55
2 100
1 2 409
outputCopy
0
8
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1399E1)
}
