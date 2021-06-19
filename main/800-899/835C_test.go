package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/835/C
// https://codeforces.com/problemset/status/835/problem/C
func TestCF835C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3 3
1 1 1
3 2 0
2 1 1 2 2
0 2 1 4 5
5 1 1 5 5
outputCopy
3
0
3
inputCopy
3 4 5
1 1 2
2 3 0
3 3 1
0 1 1 100 100
1 2 2 4 4
2 2 1 4 7
1 50 50 51 51
outputCopy
3
3
5
0`
	testutil.AssertEqualCase(t, rawText, 0, CF835C)
}
