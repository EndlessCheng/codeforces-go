package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1618/C
// https://codeforces.com/problemset/status/1618/problem/C
func TestCF1618C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 2 3 4 5
3
10 5 15
3
100 10 200
10
9 8 2 6 6 2 8 6 5 4
2
1 3
outputCopy
2
0
100
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1618C)
}
