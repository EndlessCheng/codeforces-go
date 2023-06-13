package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1840/problem/F
// https://codeforces.com/problemset/status/1840/problem/F
func TestCF1840F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 3
4
1 2 0
2 2 1
3 2 2
4 1 1
3 3
6
2 1 0
2 1 1
2 1 2
2 2 0
2 2 1
2 2 2
2 1
3
7 1 2
2 1 1
7 2 1
2 2
5
9 1 2
3 2 0
5 1 2
4 2 2
7 1 0
4 6
7
6 1 2
12 1 3
4 1 0
17 2 3
1 2 6
16 2 6
3 2 4
outputCopy
5
-1
3
6
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1840F)
}
