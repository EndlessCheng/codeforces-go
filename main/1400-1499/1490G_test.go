package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/G
// https://codeforces.com/problemset/status/1490/problem/G
func TestCF1490G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 3
1 -3 4
1 5 2
2 2
-2 0
1 2
2 2
0 1
1 2
outputCopy
0 6 2 
-1 -1 
1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1490G)
}
