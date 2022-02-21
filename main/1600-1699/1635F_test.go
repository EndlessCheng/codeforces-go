package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1635/F
// https://codeforces.com/problemset/status/1635/problem/F
func TestCF1635F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
-2 2
0 10
1 1
9 2
12 7
1 3
2 3
1 5
3 5
2 4
outputCopy
9
11
9
24
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1635F)
}
