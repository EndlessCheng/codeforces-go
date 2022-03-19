package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1374/problem/D
// https://codeforces.com/problemset/status/1374/problem/D
func TestCF1374D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 3
1 2 1 3
10 6
8 7 1 8 3 7 5 10 8 9
5 10
20 100 50 20 100500
10 25
24 24 24 24 24 24 24 24 24 24
8 8
1 2 3 4 5 6 7 8
outputCopy
6
18
0
227
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1374D)
}
