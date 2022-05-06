package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1674/problem/F
// https://codeforces.com/problemset/status/1674/problem/F
func TestCF1674F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 8
..**
.*..
*...
...*
1 3
2 3
3 1
2 3
3 4
4 3
2 3
2 2
outputCopy
3
4
4
3
4
5
5
5
inputCopy
2 5 5
*...*
*****
1 3
2 2
1 3
1 5
2 3
outputCopy
2
3
3
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1674F)
}
