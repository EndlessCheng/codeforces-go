package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1213/G
// https://codeforces.com/problemset/status/1213/problem/G
func TestCF1213G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 5
1 2 1
3 2 3
2 4 1
4 5 2
5 7 4
3 6 2
5 2 3 4 1
outputCopy
21 7 15 21 3 
inputCopy
1 2
1 2
outputCopy
0 0 
inputCopy
3 3
1 2 1
2 3 2
1 3 2
outputCopy
1 3 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1213G)
}
