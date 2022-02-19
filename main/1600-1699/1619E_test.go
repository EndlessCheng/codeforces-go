package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1619/E
// https://codeforces.com/problemset/status/1619/problem/E
func TestCF1619E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
0 1 3
7
0 1 2 3 4 3 2
4
3 0 0 0
7
4 6 2 3 5 0 5
5
4 0 1 0 4
outputCopy
1 1 0 -1 
1 1 2 2 1 0 2 6 
3 0 1 4 3 
1 0 -1 -1 -1 -1 -1 -1 
2 1 0 2 -1 -1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1619E)
}
