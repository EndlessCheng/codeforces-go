package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/B
// https://codeforces.com/problemset/status/1399/problem/B
func TestCF1399B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
3 5 6
3 2 3
5
1 2 3 4 5
5 4 3 2 1
3
1 1 1
2 2 2
6
1 1000000000 1000000000 1000000000 1000000000 1000000000
1 1 1 1 1 1
3
10 12 8
7 5 4
outputCopy
6
16
0
4999999995
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1399B)
}
