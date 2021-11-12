package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1601/B
// https://codeforces.com/problemset/status/1601/problem/B
func TestCF1601B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 2 2
1 1 0
outputCopy
2
1 0 
inputCopy
2
1 1
1 0
outputCopy
-1
inputCopy
10
0 1 2 3 5 5 6 7 8 5
9 8 7 1 5 4 3 2 0 0
outputCopy
3
9 4 0 
inputCopy
6
2 2 2 2 2 2
1 1 1 1 1 0
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1601B)
}
