package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/821/D
// https://codeforces.com/problemset/status/821/problem/D
func TestCF821D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 5
1 1
2 1
2 3
3 3
4 3
outputCopy
2
inputCopy
5 5 4
1 1
2 1
3 1
3 2
outputCopy
-1
inputCopy
2 2 4
1 1
1 2
2 1
2 2
outputCopy
0
inputCopy
5 5 4
1 1
2 2
3 3
4 4
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF821D)
}
