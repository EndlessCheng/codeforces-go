package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1728/problem/C
// https://codeforces.com/problemset/status/1728/problem/C
func TestCF1728C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
1
1000
4
1 2 3 4
3 1 4 2
3
2 9 3
1 100 9
10
75019 709259 5 611271314 9024533 81871864 9 3 6 4865
9503 2 371245467 6 7 37376159 8 364036498 52295554 169
outputCopy
2
0
2
18`
	testutil.AssertEqualCase(t, rawText, 0, CF1728C)
}
