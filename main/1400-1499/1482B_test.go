package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1482/B
// https://codeforces.com/problemset/status/1482/problem/B
func TestCF1482B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6
1 9 17 6 14 3
3
4 2 2
3
7 3 4
3
2 2 4
5
0 1000000000 0 1000000000 0
2
1 1
outputCopy
19 8
-1
-1
-1
2000000000 1000000000
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1482B)
}
