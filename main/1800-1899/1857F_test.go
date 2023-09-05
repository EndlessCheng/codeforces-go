package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1857/F
// https://codeforces.com/problemset/status/1857/problem/F
func TestCF1857F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 3 2
4
3 2
5 6
3 1
5 5
4
1 1 1 1
1
2 1
6
1 4 -2 3 3 3
3
2 -8
-1 -2
7 12
outputCopy
1 1 0 0 
6 
1 1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1857F)
}
