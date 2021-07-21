package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/507/E
// https://codeforces.com/problemset/status/507/problem/E
func TestCF507E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1
1 2 0
outputCopy
1
1 2 1
inputCopy
4 4
1 2 1
1 3 0
2 3 1
3 4 1
outputCopy
3
1 2 0
1 3 1
2 3 0
inputCopy
8 9
1 2 0
8 3 0
2 3 1
1 4 1
8 7 0
1 5 1
4 6 1
5 7 0
6 8 0
outputCopy
3
2 3 0
1 5 0
6 8 1`
	testutil.AssertEqualCase(t, rawText, 0, CF507E)
}
