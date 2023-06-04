package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1839/problem/B
// https://codeforces.com/problemset/status/1839/problem/B
func TestCF1839B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
2 2
1 6
1 10
1 13
5
3 4
3 1
2 5
3 2
3 3
6
1 2
3 4
1 4
3 4
3 5
2 3
1
1 1
outputCopy
15
14
20
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1839B)
}
