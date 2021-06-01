package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1389/B
// https://codeforces.com/problemset/status/1389/problem/B
func TestCF1389B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 4 0
1 5 4 3 2
5 4 1
1 5 4 3 2
5 4 4
10 20 30 40 50
10 7 3
4 6 8 2 9 9 7 4 10 9
outputCopy
15
19
150
56
inputCopy
1
18 11 4
11 19 18 19 19 5 14 15 17 4 10 9 8 17 9 2 15 10
outputCopy
219`
	testutil.AssertEqualCase(t, rawText, 0, CF1389B)
}
