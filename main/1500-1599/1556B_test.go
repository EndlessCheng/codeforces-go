package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1556/B
// https://codeforces.com/problemset/status/1556/problem/B
func TestCF1556B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
6 6 1
1
9
6
1 1 1 2 2 2
2
8 6
6
6 2 3 4 5 1
outputCopy
1
0
3
-1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1556B)
}
