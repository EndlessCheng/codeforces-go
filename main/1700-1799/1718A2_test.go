package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1718/A2
// https://codeforces.com/problemset/status/1718/problem/A2
func TestCF1718A2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
4
5 5 5 5
3
1 3 2
2
0 0
3
2 5 7
6
1 2 3 3 2 1
10
27 27 34 32 2 31 23 56 52 4
5
1822 1799 57 23 55
outputCopy
2
2
0
2
4
7
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1718A2)
}
