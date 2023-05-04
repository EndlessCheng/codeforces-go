package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1721/problem/C
// https://codeforces.com/problemset/status/1721/problem/C
func TestCF1721C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
2 3 5
7 11 13
1
1000
5000
4
1 2 3 4
1 2 3 4
4
10 20 30 40
22 33 33 55
outputCopy
5 4 2
11 10 8
4000
4000
0 0 0 0
0 0 0 0
12 2 3 15
23 13 3 15`
	testutil.AssertEqualCase(t, rawText, 0, CF1721C)
}
