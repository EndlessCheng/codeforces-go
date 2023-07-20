package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1772/D
// https://codeforces.com/problemset/status/1772/problem/D
func TestCF1772D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
5
5 3 3 3 5
4
5 3 4 5
8
1 2 3 4 5 6 7 8
6
10 5 4 3 2 1
3
3 3 1
3
42 43 42
2
100000000 99999999
6
29613295 52036613 75100585 78027446 81409090 73215
outputCopy
4
-1
0
42
2
-1
100000000
40741153`
	testutil.AssertEqualCase(t, rawText, 0, CF1772D)
}
