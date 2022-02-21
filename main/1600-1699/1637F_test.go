package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1637/F
// https://codeforces.com/problemset/status/1637/problem/F
func TestCF1637F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 1
1 2
2 3
outputCopy
4
inputCopy
5
1 3 3 1 3
1 3
5 4
4 3
2 3
outputCopy
7
inputCopy
2
6 1
1 2
outputCopy
12
inputCopy
6
8 8 11 3 8 9
6 5
2 5
1 3
3 6
5 4
outputCopy
25`
	testutil.AssertEqualCase(t, rawText, -1, CF1637F)
}
