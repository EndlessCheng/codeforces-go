package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1495/D
// https://codeforces.com/problemset/status/1495/problem/D
func TestCF1495D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2
2 3
3 4
1 4
outputCopy
2 1 0 1
1 2 1 0
0 1 2 1
1 0 1 2
inputCopy
8 9
1 2
1 3
1 4
2 7
3 5
3 6
4 8
2 3
3 4
outputCopy
1 0 0 0 0 0 0 0
0 2 0 0 0 0 2 0
0 0 1 0 1 1 0 0
0 0 0 2 0 0 0 2
0 0 1 0 1 1 0 0
0 0 1 0 1 1 0 0
0 2 0 0 0 0 2 0
0 0 0 2 0 0 0 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1495D)
}
