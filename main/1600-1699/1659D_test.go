package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1659/D
// https://codeforces.com/problemset/status/1659/problem/D
func TestCF1659D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
2 4 2 4
7
0 3 4 2 3 2 7
3
0 0 0
4
0 0 0 4
3
1 2 3
outputCopy
1 1 0 1 
0 1 1 0 0 0 1 
0 0 0 
0 0 0 1 
1 0 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1659D)
}
