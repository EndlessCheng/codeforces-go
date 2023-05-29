package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1830/problem/C
// https://codeforces.com/problemset/status/1830/problem/C
func TestCF1830C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
6 0
5 0
8 1
1 3
10 2
3 4
6 9
1000 3
100 701
200 801
300 901
28 5
1 12
3 20
11 14
4 9
18 19
4 3
1 4
1 4
1 4
outputCopy
5
0
0
4
839415253
140
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1830C)
}
