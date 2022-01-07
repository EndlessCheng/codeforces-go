package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1051/problem/F
// https://codeforces.com/problemset/status/1051/problem/F
func TestCF1051F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2 3
2 3 1
3 1 5
3
1 2
1 3
2 3
outputCopy
3
4
1
inputCopy
8 13
1 2 4
2 3 6
3 4 1
4 5 12
5 6 3
6 7 8
7 8 7
1 4 1
1 8 3
2 6 9
2 7 1
4 6 3
6 8 2
8
1 5
1 7
2 3
2 8
3 7
3 4
6 8
7 8
outputCopy
7
5
6
7
7
1
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1051F)
}
