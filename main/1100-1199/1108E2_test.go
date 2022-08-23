package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1108/E2
// https://codeforces.com/problemset/status/1108/problem/E2
func TestCF1108E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
2 -2 3 1 2
1 3
4 5
2 5
1 3
outputCopy
6
2
4 1 
inputCopy
5 4
2 -2 3 1 4
3 5
3 4
2 4
2 5
outputCopy
7
2
3 2 
inputCopy
1 0
1000000
outputCopy
0
0
`
	testutil.AssertEqualCase(t, rawText, 0, CF1108E2)
}
