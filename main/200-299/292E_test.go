package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/292/E
// https://codeforces.com/problemset/status/292/problem/E
func TestCF292E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 10
1 2 0 -1 3
3 1 5 -2 0
2 5
1 3 3 3
2 5
2 4
2 1
1 2 1 4
2 1
2 4
1 4 2 1
2 2
outputCopy
0
3
-1
3
2
3
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF292E)
}
