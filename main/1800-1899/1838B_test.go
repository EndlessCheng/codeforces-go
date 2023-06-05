package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1838/problem/B
// https://codeforces.com/problemset/status/1838/problem/B
func TestCF1838B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
3
1 2 3
3
1 3 2
5
1 3 2 5 4
6
4 5 6 1 2 3
9
8 7 6 3 2 1 4 5 9
10
7 10 5 1 9 8 3 2 6 4
10
8 5 10 9 2 1 3 4 6 7
10
2 3 5 7 10 1 8 6 4 9
outputCopy
2 3
1 1
5 2
1 4
9 5
8 8
6 10
5 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1838B)
}
