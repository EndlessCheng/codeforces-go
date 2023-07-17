package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1834/problem/E
// https://codeforces.com/problemset/status/1834/problem/E
func TestCF1834E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3
1 2 3
5
1 2 3 4 5
2
2 3
1
1000000000
12
1 8 4 2 3 5 7 2 9 10 11 13
12
7 2 5 4 2 1 1 2 3 11 8 9
outputCopy
4
7
1
1
16
13`
	testutil.AssertEqualCase(t, rawText, 0, CF1834E)
}
