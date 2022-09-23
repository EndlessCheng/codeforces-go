package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/522/D
// https://codeforces.com/problemset/status/522/problem/D
func TestCF522D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 1 2 3 2
1 5
2 4
3 5
outputCopy
1
-1
2
inputCopy
6 5
1 2 1 3 2 3
4 6
1 3
2 5
2 4
1 6
outputCopy
2
2
3
-1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF522D)
}
